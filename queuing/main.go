package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sirupsen/logrus"
	flag "github.com/spf13/pflag"
	goji "goji.io"
	"goji.io/pat"
)

var version = "master"

func main() {
	// helpFlag := flag.BoolP("help", "h", false, "show help")
	debugFlag := flag.BoolP("debug", "d", false, "enable debug mode")
	versionFlag := flag.BoolP("version", "v", false, "show version")
	workersFlag := flag.Int("n", 4, "The number of workers to start")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	router := goji.NewMux()

	if *debugFlag {
		router.Use(debugHandler)
	}

	router.Use(logHandler)
	router.HandleFunc(pat.Post("/work"), Collector)

	logrus.WithFields(logrus.Fields{
		"port": ":8080",
	}).Infof("starting queuing service")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}

	workers := StartDispatcher(*workersFlag)

	doneCh := make(chan struct{})

	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logrus.Warnf("could not shutdown server: %v", err)
		}

		for _, worker := range workers {
			logrus.WithFields(logrus.Fields{
				"worker": worker.ID,
			}).Infoln("sending shutdown signal")
			worker.Stop()
		}

		close(doneCh)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Printf("%v", err)
	} else {
		logrus.Printf("Server shutdown!")
	}

	<-doneCh

	logrus.WithFields(logrus.Fields{
		"port": ":8080",
	}).Print("shutdown")
}
