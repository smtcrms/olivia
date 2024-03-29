package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Message struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

var messages = SerializeMessages()

func SerializeMessages() (messages []Message) {
	err := json.Unmarshal(ReadFile("res/messages.json"), &messages)
	if err != nil {
		fmt.Println(err)
	}

	return messages
}

// GetMessage retrieves a message tag and returns a random message chose from res/messages.json
func GetMessage(tag string) string {
	for _, message := range messages {
		if message.Tag != tag {
			continue
		}

		if len(message.Messages) == 1 {
			return message.Messages[0]
		}

		return message.Messages[rand.Intn(len(message.Messages))]
	}

	return messages[0].Messages[0]
}
