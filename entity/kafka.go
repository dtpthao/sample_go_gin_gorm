package entity

type KafkaConfig struct {
	BrokerHost      string
	BrokerPort      string
	ConsumerGroupID string
}

type IKafkaService interface {
	Post(topic string, jsonMsg []byte) error
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
