package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"glintecoTask/entity"
	"log"
	"math/rand"
	"os"
)

type Message struct {
	UserId     int    `json:"user_id"`
	PostId     string `json:"post_id"`
	UserAction string `json:"user_action"`
}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		// Process the message as per your requirement here
		sess.MarkMessage(msg, "")
	}
	return nil
}

func TestProducer() {

	broker := entity.KafkaConfig{
		BrokerHost: "localhost",
		BrokerPort: "9093",
	}
	//brokers := []string{fmt.Sprintf("%s:%s", broker.Host, broker.Port)}
	//producer, err := sarama.NewSyncProducer(brokers, nil)

	producer, err := NewKafkaProducer(broker)

	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
		os.Exit(1)
	}

	// Dummy Data
	userId := [5]int{100001, 100002, 100003, 100004, 100005}
	postId := [5]string{"POST00001", "POST00002", "POST00003", "POST00004", "POST00005"}
	userAction := [5]string{"love", "like", "hate", "smile", "cry"}

	for {
		// we are going to take random data from the dummy data
		message := Message{
			UserId:     userId[rand.Intn(len(userId))],
			PostId:     postId[rand.Intn(len(postId))],
			UserAction: userAction[rand.Intn(len(userAction))],
		}

		jsonMessage, err := json.Marshal(message)

		if err != nil {
			log.Fatalln("Failed to marshal message:", err)
			os.Exit(1)
		}

		//msg := &sarama.ProducerMessage{
		//	Topic: "post-likes",
		//	Value: sarama.StringEncoder(jsonMessage),
		//}
		//
		//_, _, err = producer.SendMessage(msg)
		err = producer.Post("post-likes", jsonMessage)
		if err != nil {
			log.Fatalln("Failed to send message:", err)
			os.Exit(1)
		}
		log.Println("Message sent!")
	}
}

func TestConsumer() {

	//brokers := []string{"localhost:9093"}
	//groupID := "consumer-group"
	//
	//config := sarama.NewConfig()
	//config.Version = sarama.V2_0_0_0 // specify appropriate Kafka version
	//config.Consumer.Offsets.AutoCommit.Enable = true
	//config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	//consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	//if err != nil {
	//	log.Panicf("Error creating consumer group client: %v", err)
	//}

	broker := entity.KafkaConfig{
		BrokerHost:      "localhost",
		BrokerPort:      "9093",
		ConsumerGroupID: "consumer-group",
	}
	consumer, _ := NewKafkaConsumer(broker, []string{"post-likes"}, exampleConsumerGroupHandler{})

	//ctx := context.Background()
	for {
		//err := consumerGroup.Consume(ctx, []string{"post-likes"}, exampleConsumerGroupHandler{})
		//if err != nil {
		//	log.Panicf("Error from consumer: %v", err)
		//}
		err := consumer.Consume()
		if err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
	}
}
