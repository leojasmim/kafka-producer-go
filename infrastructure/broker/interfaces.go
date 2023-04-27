package broker

import (
	"os"

	d "github.com/leojasmim/kafka-producer-go/domain"
	k "github.com/leojasmim/kafka-producer-go/infrastructure/broker/kafka"
)

type IPublisher interface {
	SendEvent(d.IEvent) error
}

func BrokerFactory(brokerType string) IPublisher {
	if brokerType == "kafka" {
		return k.NewKafkaPublisher(
			os.Getenv("KS_KAFKA_HOSTNAME"),
			os.Getenv("KS_KAFKA_TOPICNAME"),
		)
	}
	return nil
}
