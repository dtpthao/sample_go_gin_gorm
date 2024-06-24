package kafka

import (
	"context"
	"fmt"
	"glintecoTask/entity"
	"time"

	"github.com/IBM/sarama"
)

type Consumer struct {
	Topics  []string
	Group   sarama.ConsumerGroup
	Handler sarama.ConsumerGroupHandler
	Context context.Context
}

func NewKafkaConsumer(kafkaConfig entity.KafkaConfig, topics []string, handler sarama.ConsumerGroupHandler) (*Consumer, error) {

	brokers := []string{fmt.Sprintf("%s:%s", kafkaConfig.BrokerHost, kafkaConfig.BrokerPort)}

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0 // specify appropriate Kafka version
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, kafkaConfig.ConsumerGroupID, config)
	if err != nil {
		return nil, fmt.Errorf("create consumer group error: %w", err)
	}

	return &Consumer{
		Topics:  topics,
		Group:   consumerGroup,
		Handler: handler,
		Context: context.Background(),
	}, nil
}

func (c Consumer) Consume() error {
	err := c.Group.Consume(c.Context, c.Topics, c.Handler)
	if err != nil {
		return fmt.Errorf("error from consumer: %w", err)
	}
	return nil
}
