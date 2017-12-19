package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	goji "goji.io"
	"goji.io/pat"

	"github.com/Sirupsen/logrus"
)

func main() {
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, os.Interrupt, os.Kill)

	router := goji.NewMux()
	router.Use(logHandler)
	router.HandleFunc(pat.Get("/"), homeHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logrus.Printf("server closed")
			} else {
				logrus.Fatalf("cannot start server: %v", err)
			}
		}
	}()

	<-sgn

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Warnf("could not shutdown server: %v", err)
	}

	logrus.WithFields(logrus.Fields{
		"port": ":8080",
	}).Print("shutdown")
}
