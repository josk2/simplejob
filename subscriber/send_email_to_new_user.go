package subscriber

import (
	"context"
	"errors"
	"log"
	"time"

	"simplejob/pubsub"
)

type HasUserInfo interface {
	HasUserInfo() string
}

func SendEmailToNewUser(ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Send email to new user",
		Handle: func(ctx context.Context, message *pubsub.Message) error {
			if userInfo, ok := message.Data().(HasUserInfo); ok {
				log.Println("SendEmailToNewUserHandler: Sending email to new user", userInfo)
				time.Sleep(time.Second)
				log.Println("SendEmailToNewUserHandler: Email sent to user ", userInfo)
				return nil
			} else {
				log.Println("Cannot get user data from message")
				return errors.New("Cannot get user data from message ")
			}

		},
	}

}
