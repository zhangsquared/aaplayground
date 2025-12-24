package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Note: You must use Groq-specific model IDs
// llama-3.3-70b-versatile is currently one of their best free-tier models
const GroqModel = "llama-3.3-70b-versatile"

type GroqClient struct {
	client *openai.Client
}

func InitGroqClient() (*GroqClient, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY environment variable is not set")
	}

	// Groq uses the exact same SDK but requires a custom BaseURL
	c := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL("https://api.groq.com/openai/v1"),
	)

	return &GroqClient{client: &c}, nil
}

func (c *GroqClient) Generate(ctx context.Context, prompt string) (string, error) {

	chat, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModel(GroqModel),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})

	if err != nil {
		return "", fmt.Errorf("groq api error: %w", err)
	}

	if len(chat.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from Groq")
	}

	return chat.Choices[0].Message.Content, nil
}
