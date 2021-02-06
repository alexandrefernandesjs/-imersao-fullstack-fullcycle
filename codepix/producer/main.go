package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	producer := NewKafkaProducer()

	topic := os.Getenv("kafkaTopicDesafio")
	msg := "Higurashi no Naku Koro ni é uma excelente Visual Novel e você deveria lê-la."
	deliveryChan := make(chan ckafka.Event)

	err := Publish(msg, topic, producer, deliveryChan)
	if err != nil{
		panic(err)
	}

	DeliveryReport(deliveryChan)
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")
	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: []byte(msg),
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed: ", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to: ", ev.TopicPartition)
			}
		}
	}
}