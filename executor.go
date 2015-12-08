package RxGo

import (
	"time"
)

type Task struct {
	Call      runnable
	InitDelay time.Duration
	Period    time.Duration
}

type Executor struct {
	ID      int
	Pool    chan Task
	Running bool
}

func (e *Executor) Start() {
	e.Running = true
	go func() {
		select {
		case t, more := <-e.Pool:
			if more {
				time.Sleep(t.InitDelay)
				t.Call()
			}

			if t.Period != 0 {
				go func() {
					for e.Running {
						time.Sleep(t.Period)
						t.Call()
					}
				}()
			}
		}
	}()
}

func (e *Executor) Stop() {
	close(e.Pool)
	e.Running = false
}

func (e *Executor) Submit(t Task) {
	e.Pool <- t
}
