package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simplejob/pubsub"
	"simplejob/pubsub/pblocal"
)

type Stock struct {
	Code  string
	Price float64
}

func main() {
	var CEOTopic pubsub.Topic = "CEO"
	var pbLocal = pblocal.NewPubSub()

	ctx := context.Background()

	prices := []Stock{
		{
			Code:  "CEO",
			Price: 15.5,
		},
		{
			Code:  "CEO",
			Price: 15.8,
		},
		{
			Code:  "CEO",
			Price: 16.2,
		},
		{
			Code:  "CEO",
			Price: 17.4,
		},
		{
			Code:  "CEO",
			Price: 17.1,
		},
	}
	go func() {
		for _, price := range prices {
			pbLocal.Publish(ctx, CEOTopic, pubsub.NewMessage(price))
			time.Sleep(time.Second)
		}
	}()

	ceoSub1, ceoClose1 := pbLocal.Subscribe(ctx, CEOTopic)
	ceoSub2, _ := pbLocal.Subscribe(ctx, CEOTopic)

	go func() {
		//rut data tu sub1
		for {
			data := <-ceoSub1
			log.Println("Sub1 get data: ", data.Data())
			time.Sleep(350 * time.Millisecond)
		}
	}()

	go func() {
		//rut data tu sub2
		for {
			data := <-ceoSub2
			log.Println("Sub2 get data: ", data.Data())
			time.Sleep(50 * time.Millisecond)
		}
	}()

	//close1 and republish new message
	time.Sleep(3 * time.Second)
	ceoClose1()
	//pbLocal.Publish(ctx, CEOTopic, pubsub.NewMessage(Stock{Code: "CEO", Price: 18.5}))
	G(ctx)
}
func G(ctx context.Context) {
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
