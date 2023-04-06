package asyncjob

import "time"

type jobConfig struct {
	Retries       []time.Duration
	MaxTimeout    time.Duration
	isConcurrency bool
}
