package entity

import "github.com/IBM/sarama"

type KafkaConfig struct {
	BrokerHost      string
	BrokerPort      string
	ConsumerGroupID string
}

type IKafkaService interface {
	Init(topics []string, consumerHandler sarama.ConsumerGroupHandler) error
	Post(topic string, msg any) error
	Listen()
}

type KafkaContractDeleteMessage struct {
	ActionUserIsAdmin bool   `json:"is_admin"`
	ActionUserUuid    string `json:"user_uuid"`
	ContractUuid      string `json:"contract_uuid"`
}

type KafkaUserDeleteMessage struct {
	ActionUuid string `json:"action_uuid"`
	TargetUuid string `json:"target_uuid"`
}
