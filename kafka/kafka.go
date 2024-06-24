package kafka

import (
	"github.com/IBM/sarama"
	"glintecoTask/entity"
)

type Service struct {
	Producer Producer
	Consumer Consumer
}

func NewKafkaService(config entity.KafkaConfig, topics []string, consumerHandler sarama.ConsumerGroupHandler) (*Service, error) {
	producer, err := NewKafkaProducer(config)
	if err != nil {
		return nil, err
	}

	consumer, err := NewKafkaConsumer(config, topics, consumerHandler)
	if err != nil {
		return nil, err
	}
	return &Service{*producer, *consumer}, nil
}

func (s Service) Post(topic string, jsonMsg []byte) error {
	return s.Producer.Post(topic, jsonMsg)
}

func (s Service) Consume() error {
	return s.Consume()
}
