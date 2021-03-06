package schedulers

import (
	"runtime"
	"time"
)

type Runnable func()

type Scheduler interface {
	Start()
	Stop()
	Schedule(run Runnable)
	ScheduleAt(run Runnable, delay time.Duration)
}

var (
	Computation Scheduler
	IO          Scheduler
)

func init() {
	Computation = newEventLoopScheduler(maxParallelism())
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCpu := runtime.NumCPU()
	if maxProcs < numCpu {
		return maxProcs
	}
	return numCpu
}
