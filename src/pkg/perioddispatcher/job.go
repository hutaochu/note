package perioddispatcher

import (
	"sync"
	"time"
)

type BaseJob struct {
	// Interval the frequency to execute of  this job
	Interval             time.Duration
	mu                   sync.Mutex
	activeTaskCount      int64
	nextAllowedExecution int64
}

// Name get job name, default ""
func (j *BaseJob) Name() string {
	return ""
}

//
func (j *BaseJob) GetInterval() time.Duration {
	return j.Interval
}
