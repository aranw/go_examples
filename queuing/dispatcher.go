package main

import "github.com/Sirupsen/logrus"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworkers int) []*Worker {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)
	workers := make([]*Worker, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		logrus.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
		workers[i] = worker
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				logrus.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					logrus.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()

	return workers
}
