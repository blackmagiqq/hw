package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var isIngoringErrors bool

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksChannel := make(chan Task)
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	errsCount := 0

	if m == 0 {
		isIngoringErrors = true
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(tasksChannel, &mu, &wg, &errsCount)
	}

	for _, task := range tasks {
		mu.RLock()
		isOverLimit := errsCount >= m
		mu.RUnlock()
		if isOverLimit && !isIngoringErrors {
			break
		}
		tasksChannel <- task
	}
	close(tasksChannel)

	wg.Wait()

	if errsCount >= m && !isIngoringErrors {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(tasks <-chan Task, mu *sync.RWMutex, wg *sync.WaitGroup, countErrs *int) {
	defer wg.Done()
	for task := range tasks {
		if err := task(); err != nil {
			mu.Lock()
			*countErrs++
			mu.Unlock()
		}
	}
}
