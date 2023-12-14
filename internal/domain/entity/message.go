package entity

import (
	"CHAT_SERVICE_API/util"
	"errors"
	"github.com/google/uuid"
	"github.com/pkoukk/tiktoken-go"
	"time"
)

type Message struct {
	ID        string
	Content   string
	Role      int
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role int, content string, model *Model) (*Message, error) {
	tkm, err := tiktoken.EncodingForModel(model.GetName())
	if err != nil {
		return nil, err
	}

	tokens := tkm.Encode(content, nil, nil)

	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Model:     model,
		Tokens:    len(tokens),
		CreatedAt: time.Now(),
	}
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != util.USER && m.Role != util.SYSTEM && m.Role != util.ASSISTANT {
		return errors.New("invalid role")
	}
	if m.Content == "" {
		return errors.New("content is empty")
	}
	return nil
}
