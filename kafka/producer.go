package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"glintecoTask/entity"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(config entity.KafkaConfig) (*Producer, error) {
	brokers := []string{fmt.Sprintf("%s:%s", config.BrokerHost, config.BrokerPort)}

	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start Sarama producer: %w", err)
	}

	return &Producer{Producer: producer}, nil
}

func (p Producer) Post(topic string, jsonMsg []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonMsg),
	}

	_, _, err := p.Producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
