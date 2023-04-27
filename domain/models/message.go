package models

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	ID        string    `json:"Id"`
	Payload   string    `json:"Payload"`
	CreatedAt time.Time `json:"Created_at"`
	Source    string    `json:"Source"`
}

func NewMessage(p string) *Message {
	return &Message{
		ID:        GenerateID("msg"),
		Payload:   p,
		CreatedAt: time.Now(),
		Source:    "golang",
	}
}

func (m Message) GetEvent() (string, []byte) {
	jstring, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Fail convert data", err)
	}
	return m.ID, []byte(jstring)
}
