package asyncjob

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestJob_Status(t *testing.T) {

	j := NewJob(func(ctx context.Context) error {
		fmt.Println("Test job")
		return errors.New("error something")
	})

	err := j.Execute(context.Background())
	if err != nil {
		fmt.Println("run va loi", err)

	}
}
