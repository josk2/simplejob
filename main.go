package main

import (
	"context"
	"errors"
	"fmt"

	"simplejob/asyncjob"
)

func main() {
	j := asyncjob.NewJob(func(ctx context.Context) error {
		fmt.Println("Test job")
		return errors.New("error something")
	})

	ctx := context.Background()
	err := j.Execute(ctx)
	if err != nil {
		fmt.Println("run va loi", err)
		j.Retry(ctx)
	}
}
