package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	taskCh := make(chan Task)
	var errC int32
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg, taskCh, &errC)
	}

	mInt32 := int32(m)

	for _, task := range tasks {
		if atomic.LoadInt32(&errC) >= mInt32 {
			break
		}
		taskCh <- task
	}
	close(taskCh)

	wg.Wait()

	if atomic.LoadInt32(&errC) > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(wg *sync.WaitGroup, ch <-chan Task, errC *int32) {
	defer wg.Done()
	for task := range ch {
		if err := task(); err != nil {
			atomic.AddInt32(errC, 1)
		}
	}
}
