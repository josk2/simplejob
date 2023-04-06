package asyncjob

import "context"

type Group struct {
	Jobs       []job
	isCurrency bool
}

func NewGroup(isConcurency bool, jobs ...job) *Group {
	return &Group{
		isCurrency: isConcurency,
		Jobs:       jobs,
	}
}
func (g Group) Run(ctx context.Context) error {
	for _, job := range g.Jobs {
		job.Execute(ctx)
	}
	return nil
}

func (g *Group) AddJob(jobs ...job) {
	for _, j := range jobs {
		g.Jobs = append(g.Jobs, j)
	}
}
