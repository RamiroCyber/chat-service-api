package entity

import (
	"CHAT_SERVICE_API/util"
	"errors"
)

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID             string
	UserID         string
	InitialMessage *Message
	Messages       []*Message
	ErasedMessages []*Message
	Status         int
	TokenUsage     int
	Config         *ChatConfig
}

func (c *Chat) AddMessage(message *Message) error {
	if c.Status == util.ENDED {
		errors.New("chat is ended. no more messages allowed")
	}
	for {
		if c.Config.Model.GetMaxTokens() >= message.GetQtdTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, message)
			c.updateTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.updateTokenUsage()
	}
	return nil
}

func (c *Chat) Close() {
	c.Status = util.ENDED
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) updateTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].GetQtdTokens()
	}
}
