package main

import (
	"context"
	"errors"
	"log"
	"time"

	"simplejob/asyncjob"
)

func main() {
	ctx := context.Background()
	j0 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Testing job 0")
		return errors.New("error something")
	})
	j0.SetRetryTime([]time.Duration{time.Second * 2, time.Second * 5})

	if e := j0.Execute(ctx); e != nil {
		log.Printf("Job 0 error when execute: %v", e)
		for {
			if err := j0.Retry(ctx); err == nil {
				break
			}

			if j0.State() == asyncjob.StateRetryFailed {
				break
			}
			log.Printf("Job 0 error when retry: %v", e)
		}
	}

	//j1 := asyncjob.NewJob(func(ctx context.Context) error {
	//	log.Println("Test job A")
	//	log.Println("job A sleeping 1s")
	//	return errors.New("A: error something")
	//})
	//j1.SetRetryTime([]time.Duration{time.Second * 2, time.Second * 5})
	//
	//j2 := asyncjob.NewJob(func(ctx context.Context) error {
	//	log.Println("Test job B")
	//
	//	time.Sleep(3 * time.Second)
	//	log.Println("job B sleeping 3s")
	//	return nil
	//})
	//j2.SetRetryTime([]time.Duration{time.Second * 3, time.Second * 3})
	//
	//j3 := asyncjob.NewJob(func(ctx context.Context) error {
	//	log.Println("Test job C")
	//
	//	time.Sleep(3 * time.Second)
	//	log.Println("job C sleeping 5s")
	//	return errors.New("C: error something")
	//})
	//j3.SetRetryTime([]time.Duration{time.Second, time.Second * 3})
	//
	//group := asyncjob.NewGroup(true, j1, j2, j3)
	//group.Run(ctx)

	time.Sleep(10 * time.Second)
}
