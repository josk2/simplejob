package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	data      interface{}
	channel   Topic
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()
	return &Message{
		id:        fmt.Sprintf("%v", now.UnixMicro()),
		data:      data,
		createdAt: now,
	}
}

func (m *Message) Channel() Topic {
	return m.channel
}

func (m *Message) SetChanel(topic Topic) *Message {
	m.channel = topic
	return m
}

func (m *Message) String() string {
	return fmt.Sprintf("Message %v", m.id)
}

func (m *Message) Data() interface{} {
	return m.data
}
