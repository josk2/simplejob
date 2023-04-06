package asyncjob

import (
	"context"
	"errors"
	"time"
)

const DefaultMaxTimeout = 10 * time.Second

const (
	StateInit = iota
	StateRunning
	StateFailed
	StateRetryFailed
	StateCompleted
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	SetRetryTime(retry []time.Duration)
	State() JobState
}

var defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 5}

type JobHandler func(ctx context.Context) error

type JobState int

func (js JobState) State() string {
	return []string{"Init", "Running", "Retry Failed", "Completed"}[js]
}

type job struct {
	Handler    JobHandler
	config     *jobConfig
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	return &job{
		Handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		config: &jobConfig{
			Retries:    defaultRetryTime,
			MaxTimeout: 10 * time.Second,
		},
	}
}

func (j *job) SetRetryTime(retry []time.Duration) *job {
	if len(retry) == 0 {
		return j
	}
	j.config.Retries = retry
	return j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning
	err := j.Handler(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1

	if j.retryIndex == len(j.config.Retries) {
		j.state = StateRetryFailed
		return errors.New("Retry quota has been reached")
	}

	retryTime := j.config.Retries[j.retryIndex]
	time.Sleep(retryTime)

	return j.Execute(ctx)
}
