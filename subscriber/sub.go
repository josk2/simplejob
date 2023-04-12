package subscriber

import (
	"context"
	"log"

	"simplejob/asyncjob"
	"simplejob/common"
	"simplejob/pubsub"

	"gorm.io/gorm"
)

type consumerJob struct {
	Title  string
	Handle func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	db *gorm.DB
	ps pubsub.PubSub
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func NewConsumerEngine(db *gorm.DB, ps pubsub.PubSub) *consumerEngine {
	return &consumerEngine{
		db: db,
		ps: ps,
	}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicNewUserRegistered,
		true,
		SendEmailToNewUser(context.Background()),
		NotifyAdminAfterSendEmail(context.Background()),
	)

	return nil
}

func (engine *consumerEngine) startSubTopic(
	topic pubsub.Topic,
	isConcurrent bool,
	consumerJobs ...consumerJob,
) error {
	channel, _ := engine.ps.Subscribe(context.Background(), topic)

	//convert consumer job to async job
	convertJobHandler := func(cjob consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			//do something
			log.Println("running job for: ", cjob.Title, ".Data", message.Data())
			return cjob.Handle(ctx, message)
		}
	}

	go func() {
		for {
			message := <-channel
			//get list job
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))
			for i, _ := range consumerJobs {
				jobHdlArr[i] = asyncjob.NewJob(convertJobHandler(consumerJobs[i], message))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
