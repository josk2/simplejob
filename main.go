package main

import (
	"context"
	"log"
	"time"

	"simplejob/asyncjob"
)

func main() {
	ctx := context.Background()

	jobExample(ctx)

}

func jobExample(ctx context.Context) {
	j1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Test job A")
		log.Println("job A sleeping 1s")
		time.Sleep(time.Second)
		return nil
		//return errors.New("A: error something")
	})
	j1.SetRetryTime([]time.Duration{time.Second * 2, time.Second * 5})

	j2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Test job B")

		time.Sleep(3 * time.Second)
		log.Println("job B sleeping 3s")
		return nil
	})
	j2.SetRetryTime([]time.Duration{time.Second * 3, time.Second * 3})

	j3 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Test job C")

		time.Sleep(1 * time.Second)
		log.Println("job C sleeping 1s")
		return nil
	})
	j3.SetRetryTime([]time.Duration{time.Second, time.Second * 3})

	group := asyncjob.NewGroup(true, j1, j2, j3)
	group.Run(ctx)
}
