package asyncjob

import (
	"context"
	"time"
)

const DefaultMaxTimeout = 10 * time.Second

const (
	StateInit = iota
	StateRunning
	StateRetryFailed
	StateCompleted
)

var defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 5}

type JobHandler func(ctx context.Context) error

type job struct {
	retryTime   []time.Duration
	Handler     func() error
	statusIndex int32
}

func NewJob(handler func() error) *job {
	return &job{
		retryTime:   defaultRetryTime,
		Handler:     handler,
		statusIndex: StatusInit,
	}
}

type JobState int

func (j JobState) State() string {
	return []string{"Init", "Running", "Retry Failed", "Completed"}[j]
}
