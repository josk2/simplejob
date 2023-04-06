package main

import (
	"context"
	"errors"
	"log"
	"time"

	"simplejob/asyncjob"
)

func main() {
	j1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Test job A")
		time.Sleep(1 * time.Second)
		return errors.New("A: error something")
	})
	j2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Test job B")
		time.Sleep(3 * time.Second)
		return nil
	})

	ctx := context.Background()
	group := asyncjob.NewGroup(true, j1, j2)
	group.Run(ctx)
}
