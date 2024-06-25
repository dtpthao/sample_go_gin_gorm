package test

import "github.com/IBM/sarama"

type MockKafka struct{}

func (m MockKafka) Init(topics []string, consumerHandler sarama.ConsumerGroupHandler) error {
	// ok fine
	return nil
}

func (m MockKafka) Post(topic string, msg any) error {
	// pushed success
	return nil
}

func (m MockKafka) Listen() {
	// actively listening
}
