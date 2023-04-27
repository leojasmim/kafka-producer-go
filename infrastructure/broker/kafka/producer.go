package broker_kafka

import (
	"context"
	"log"
	"time"

	"github.com/leojasmim/kafka-producer-go/domain"
	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Host       string
	TopicName  string
	connection *kafka.Conn
}

func NewKafkaPublisher(host, topic string) *KafkaPublisher {
	leaderPartition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", host, topic, leaderPartition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return nil
	}

	return &KafkaPublisher{
		Host:       host,
		TopicName:  topic,
		connection: conn,
	}
}

func (k *KafkaPublisher) SendEvent(msg domain.IEvent) error {
	defer k.closeConnection()

	key, value := msg.GetEvent()

	k.connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := k.connection.WriteMessages(
		kafka.Message{Key: []byte(key), Value: value},
	)
	return err
}

func (k *KafkaPublisher) closeConnection() {
	if err := k.connection.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
