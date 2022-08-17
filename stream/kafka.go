package stream

import "github.com/Shopify/sarama"

type Config struct {
	Brokers []string
}

func NewKafkaClient(cfg Config) (sarama.Client, error) {
	return sarama.NewClient(cfg.Brokers, nil)
}
