package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goji "goji.io"
	"goji.io/pat"
)

func main() {
	router := goji.NewMux()
	router.Use(logHandler)
	router.HandleFunc(pat.Get("/"), homeHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

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
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Printf("Server shutdown!")
	}
}
