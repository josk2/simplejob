package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simplejob/asyncjob"
	"simplejob/common"
	"simplejob/model"
	"simplejob/pubsub"
	"simplejob/subscriber"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	//jobExample(ctx)

	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println(&db)

	ps := pubsub.NewPubSub()
	consumerEngine := subscriber.NewConsumerEngine(db, ps)
	consumerEngine.Start()

	//fake publisher
	fakePublisher(ctx, ps)

	//graceful shutdown
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGTERM, syscall.SIGINT)

	//wait signal to be put
	sign := <-signChan
	log.Println("Get signal: ", sign)
	//shutdown services
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Println("Server stopped, bye bye!!!")

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

func fakePublisher(ctx context.Context, ps pubsub.PubSub) {
	users := []model.User{
		{Id: 1,
			Name:   "Anton",
			Gender: model.UserGenderMale},
		//{Id: 2,
		//	Name:   "Andy",
		//	Gender: model.UserGenderMale},
		//{Id: 1,
		//	Name:   "Anna",
		//	Gender: model.UserGenderFemale},
	}

	for _, user := range users {
		if err := ps.Publish(ctx, common.TopicNewUserRegistered, pubsub.NewMessage(user)); err != nil {
			log.Println("Publish failed", err)
		}
		time.Sleep(time.Second)
	}
}
