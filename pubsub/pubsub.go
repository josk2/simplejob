package pubsub

import "context"

type Topic string

type PubSub interface {
	Publish(ctx context.Context, channel Topic, message *Message) error
	Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, close func())
}
