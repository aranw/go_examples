package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/felixge/httpsnoop"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func logHandler(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(inner, w, r)

		logrus.WithFields(logrus.Fields{
			"duration": m.Duration,
			"path":     r.URL.Path,
			"method":   r.Method,
			"status":   m.Code,
			"written":  m.Written,
		}).Info("request completed")
	}
	return http.HandlerFunc(mw)
}
