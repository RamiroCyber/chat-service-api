package entity

import (
	"errors"
	"github.com/google/uuid"
)

const (
	ENDED = iota + 1
	ACTIVE
)

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

func NewChat(userID string, initialMessage *Message, config *ChatConfig) (*Chat, []string) {
	chat := &Chat{
		ID:             uuid.New().String(),
		UserID:         userID,
		InitialMessage: initialMessage,
		Config:         config,
		Status:         ACTIVE,
		TokenUsage:     0,
	}
	chat.AddMessage(initialMessage)
	if err := chat.Validate(); err != nil {
		return nil, err
	}
	return chat, nil
}

func (c *Chat) Validate() (err []string) {
	if c.UserID == "" {
		err = append(err, "user id is empty;")
	}

	if c.Status != ACTIVE && c.Status != ENDED {
		err = append(err, "invalid status;")
	}

	if c.Config.Temperature < 0 || c.Config.Temperature > 2 {
		err = append(err, "invalid temperature;")
	}
	return err
}

func (c *Chat) AddMessage(message *Message) error {
	if c.Status == ENDED {
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
	c.Status = ENDED
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
