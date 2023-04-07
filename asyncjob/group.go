package asyncjob

import (
	"context"
	"log"
	"sync"
)

type Group struct {
	Jobs       []job
	isCurrency bool
}

var wg *sync.WaitGroup

func NewGroup(isConcurency bool, jobs ...job) *Group {
	return &Group{
		isCurrency: isConcurency,
		Jobs:       jobs,
	}
}
func (g Group) Run(ctx context.Context) error {

	var err error
	for _, job := range g.Jobs {
		jerr := job.Execute(ctx)
		if jerr != nil {
			log.Printf("job err: %v", jerr)
			err = jerr
		}
	}

	return err
}

func (g *Group) AddJob(jobs ...job) {
	for _, j := range jobs {
		g.Jobs = append(g.Jobs, j)
	}
}
