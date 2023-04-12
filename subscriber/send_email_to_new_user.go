package subscriber

import (
	"context"
	"errors"
	"log"
	"time"

	"simplejob/model"
	"simplejob/pubsub"
)

type HasUserInfo interface {
	HasUserInfo() string
}

func SendEmailToNewUser(ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Send email to new user",
		Handle: func(ctx context.Context, message *pubsub.Message) error {
			if userInfo, ok := message.Data().(model.User); ok {
				log.Println("SendEmailToNewUserHandler: Sending email to new user", userInfo.Name)
				time.Sleep(time.Second)
				log.Println("SendEmailToNewUserHandler: Email sent to user", userInfo.Name)
				return nil
			} else {
				log.Println("SendEmailToNewUserHandler Cannot get user data from message", userInfo)
				return errors.New("SendEmailToNewUserHandler Cannot get user data from message ")
			}

		},
	}

}

func NotifyAdminAfterSendEmail(ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Notify admin after send email to new user",
		Handle: func(ctx context.Context, message *pubsub.Message) error {
			if userInfo, ok := message.Data().(model.User); ok {
				log.Println("NotifyAdminAfterSendEmailHandler: calling GCM", userInfo.Name)
				time.Sleep(time.Second)
				log.Println("NotifyAdminAfterSendEmailHandler: Noticed", userInfo.Name)
				return nil
			} else {
				log.Println("NotifyAdminAfterSendEmailHandler Cannot get user data from message", userInfo)
				return errors.New("NotifyAdminAfterSendEmailHandler Cannot get user data from message ")
			}

		},
	}

}
