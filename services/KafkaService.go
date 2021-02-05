package services

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"fmt"
	"encoding/json"

)

func SendImagesToKafka(base64file string){

	//Init to get request data
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")})
	if err != nil {
		panic(err)

	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	aiMessage := &AiMessage{ImageData: base64file}
	message, err := json.Marshal(aiMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Produce messages to topic (asynchronously)
	topic := os.Getenv("KAFKA_TOPIC_IMAGE")
	_ = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	p.Flush(36000)

}


type AiMessage struct {
	ImageData string `json:"imagedata"`
}
