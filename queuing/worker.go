package main

import (
	"time"

	"github.com/Sirupsen/logrus"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) *Worker {
	// Create, and return the worker.
	return &Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.
				logrus.WithFields(logrus.Fields{
					"worker":  w.ID,
					"delayed": work.Delay.Seconds(),
				}).Println("Received work request")

				time.Sleep(work.Delay)
				logrus.WithFields(logrus.Fields{
					"worker": w.ID,
				}).Printf("Hello, %s!\n", work.Name)

			case <-w.QuitChan:
				// We have been asked to stop.
				logrus.WithFields(logrus.Fields{
					"worker": w.ID,
				}).Println("stopping")
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	logrus.WithFields(logrus.Fields{
		"worker": w.ID,
	}).Println("Received stop signal")
	w.QuitChan <- true
}
