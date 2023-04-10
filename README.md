## Simple Job - Job Manager

### Job:
Các job:
- retry, retry sau một thời gian nhất định
- có thể config bật/tắt concurrency


### Job Manager (Group)
- quản lý các job
- Có thể cài đặt cho các job chạy đồng thời


## Simple Pubsub
- Message:
```golang
type Topic string

type Message struct {
	id        string
	data      interface{}
	channel   Topic
	createdAt time.Time
}
```

- Pubsub interface
```go
type PubSub interface {
	Publish(ctx context.Context, channel Topic, message *Message) error
	Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, close func())
}
```