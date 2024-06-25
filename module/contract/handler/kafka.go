package handler

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"sample-go-server/entity"
)

type KafkaDeleteContractHandler struct {
	uc entity.IContractUseCase
}

func NewKafkaContractHandler(uc entity.IContractUseCase) sarama.ConsumerGroupHandler {
	return KafkaDeleteContractHandler{uc}
}

func (KafkaDeleteContractHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (KafkaDeleteContractHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h KafkaDeleteContractHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var deleteMsg entity.KafkaContractDeleteMessage
		err := json.Unmarshal(msg.Value, &deleteMsg)
		if err != nil {
			sess.MarkMessage(msg, "")
			return err
		}

		if deleteMsg.ActionUserIsAdmin {
			err = h.uc.Delete(deleteMsg.ContractUuid)
			if err != nil {
				sess.MarkMessage(msg, "")
				return err
			}
		} else {
			err = h.uc.DeleteByUser(deleteMsg.ContractUuid, deleteMsg.ActionUserUuid)
			if err != nil {
				sess.MarkMessage(msg, "")
				return err
			}
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}
