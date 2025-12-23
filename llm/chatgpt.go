package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type ChatGPTClient struct {
	client *openai.Client
}

func InitChatGPTClient() (*ChatGPTClient, error) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	c := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	return &ChatGPTClient{client: &c}, nil
}

func (c *ChatGPTClient) Generate(ctx context.Context, prompt string) (string, error) {
	chat, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})

	if err != nil {
		return "", err
	}

	if len(chat.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from ChatGPT")
	}

	return chat.Choices[0].Message.Content, nil
}
