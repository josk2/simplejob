package pubsub

import (
	"context"
	"log"
	"sync"
)

const MessageBuffer = 10000

type engine struct {
	messageQueue chan *Message
	mapChannel   map[Topic][]chan *Message
	locker       *sync.RWMutex
}

func NewPubSub() *engine {
	ps := &engine{
		messageQueue: make(chan *Message, MessageBuffer),
		mapChannel:   make(map[Topic][]chan *Message),
		locker:       new(sync.RWMutex),
	}

	//run in here
	err := ps.run()
	if err != nil {
		log.Fatal(err)
	}
	return ps
}

func (e *engine) Publish(ctx context.Context, topic Topic, message *Message) error {
	message.SetChanel(topic)

	//send message to queue
	go func() {
		e.messageQueue <- message
	}()

	return nil
}

func (e *engine) Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, close func()) {
	c := make(chan *Message)

	e.locker.Lock()
	if vals, ok := e.mapChannel[topic]; ok {
		vals = append(e.mapChannel[topic], c)
		e.mapChannel[topic] = vals
	} else {
		e.mapChannel[topic] = []chan *Message{c}
	}
	e.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		//remove chanel from []chan *Message
		if chans, ok := e.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					//save to pubsub channels
					e.locker.Lock()
					e.mapChannel[topic] = chans
					e.locker.Unlock()
					break
				}
			}
		}
	}
}

func (e *engine) run() error {
	log.Println("started pubsub")
	go func() {
		for {
			mess := <-e.messageQueue
			log.Println("Message dequeue:", mess)
			if subs, ok := e.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
