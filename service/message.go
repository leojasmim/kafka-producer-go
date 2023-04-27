package service

import (
	"github.com/leojasmim/kafka-producer-go/domain/models"
	"github.com/leojasmim/kafka-producer-go/infrastructure/broker"
)

type MessageService struct {
	publisher broker.IPublisher
}

func NewMessageService() *MessageService {
	return &MessageService{
		publisher: broker.BrokerFactory("kafka"),
	}
}

func (s *MessageService) SendMessage(msg models.Message) error {
	return s.publisher.SendEvent(msg)
}
