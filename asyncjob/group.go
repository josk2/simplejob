package asyncjob

import (
	"context"
	"log"
	"sync"
)

type group struct {
	Jobs       []job
	isCurrency bool
	wg         *sync.WaitGroup
}

func NewGroup(isConcurency bool, jobs ...job) *group {
	return &group{
		isCurrency: isConcurency,
		Jobs:       jobs,
		wg:         new(sync.WaitGroup),
	}
}
func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.Jobs))

	errChan := make(chan error, len(g.Jobs))

	//run job
	for i, _ := range g.Jobs {
		if g.isCurrency {
			//run job concurrency
			go func(aj *job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(&g.Jobs[i])
			continue
		}

		//run sequence
		job := g.Jobs[i]
		errChan <- g.runJob(ctx, &job)
		g.wg.Done()
	}

	var err error

	for i := 0; i < len(g.Jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}

	g.wg.Wait()
	return err
}

func (g *group) runJob(ctx context.Context, job *job) error {

	var e error
	if e = job.Execute(ctx); e != nil {
		log.Printf("Job error when execute: %v", e)
		for {
			if e = job.Retry(ctx); e == nil {
				break
			}

			if job.State() == StateRetryFailed {
				break
			}
		}
		return e
	}

	return nil
}

func (g *group) AddJob(jobs ...job) {
	for _, j := range jobs {
		g.Jobs = append(g.Jobs, j)
	}
}
