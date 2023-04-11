package pblocal

import (
	"context"
	"log"
	"sync"

	ps "simplejob/pubsub"
)

const MessageBuffer = 10000

type localPubSub struct {
	messageQueue chan *ps.Message
	mapChannel   map[ps.Topic][]chan *ps.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	ps := &localPubSub{
		messageQueue: make(chan *ps.Message, MessageBuffer),
		mapChannel:   make(map[ps.Topic][]chan *ps.Message),
		locker:       new(sync.RWMutex),
	}

	//run in here
	err := ps.run()
	if err != nil {
		log.Fatal(err)
	}
	return ps
}

func (l *localPubSub) Publish(ctx context.Context, topic ps.Topic, message *ps.Message) error {
	message.SetChanel(topic)

	//send message to queue
	go func() {
		l.messageQueue <- message
	}()

	return nil
}

func (l *localPubSub) Subscribe(ctx context.Context, topic ps.Topic) (ch <-chan *ps.Message, close func()) {
	c := make(chan *ps.Message)

	l.locker.Lock()
	if vals, ok := l.mapChannel[topic]; ok {
		vals = append(l.mapChannel[topic], c)
		l.mapChannel[topic] = vals
	} else {
		l.mapChannel[topic] = []chan *ps.Message{c}
	}
	l.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		//remove chanel from []chan *Message
		if chans, ok := l.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					//save to pubsub channels
					l.locker.Lock()
					l.mapChannel[topic] = chans
					l.locker.Unlock()
					break
				}
			}
		}
	}
}

func (l *localPubSub) run() error {
	log.Println("started pubsub")
	go func() {
		for {
			mess := <-l.messageQueue
			log.Println("Message dequeue:", mess)
			if subs, ok := l.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *ps.Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
