package stream

import (
	"context"
	"sync"

	// 3rd party
	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)

type topicName = string

type EventPublisher struct {
	client    sarama.Client
	producers map[topicName]sarama.AsyncProducer
	sync.RWMutex
}

func NewEventPublisher(client sarama.Client) *EventPublisher {
	return &EventPublisher{
		client:    client,
		producers: make(map[topicName]sarama.AsyncProducer, 0),
	}
}

func (ep *EventPublisher) Publish(ctx context.Context, topic topicName, key string, pbMessage proto.Message) error {
	pr, _ := ep.getProducer(ctx, topic)

	msg, err := proto.Marshal(pbMessage)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
		Key:   sarama.StringEncoder(key),
	}

	pr.Input() <- message

	return nil
}

func (ep *EventPublisher) getProducer(_ context.Context, topic topicName) (sarama.AsyncProducer, error) {
	ep.RLock()
	if p, ok := ep.producers[topic]; ok {
		ep.RUnlock()
		return p, nil
	}
	ep.RUnlock()

	ep.Lock()
	defer ep.Unlock()

	p, err := sarama.NewAsyncProducerFromClient(ep.client)
	if err != nil {
		return nil, err
	}
	ep.producers[topic] = p

	return p, err
}

func (ep *EventPublisher) Close() {
	for _, p := range ep.producers {
		_ = p.Close()
	}
}
