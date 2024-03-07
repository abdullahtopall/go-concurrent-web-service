package utils

import (
	"sync"
)

type Worker struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func NewWorker(size int) *Worker {
	worker := &Worker{
		tasks: make(chan func()),
	}

	for i := 0; i < size; i++ {
		go worker.worker()
	}

	return worker
}

func (w *Worker) worker() {
	for task := range w.tasks {
		task()
		w.wg.Done()
	}
}

func (w *Worker) AddTask(task func()) {
	w.wg.Add(1)
	w.tasks <- task
}

func (w *Worker) Wait() {
	w.wg.Wait()
	close(w.tasks)
}
