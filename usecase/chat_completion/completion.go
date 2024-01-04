package chat_completion

import (
	"CHAT_SERVICE_API/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}
