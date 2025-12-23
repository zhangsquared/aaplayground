package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type ClaudeClient struct {
	client *anthropic.Client
}

func InitClaudeClient() (*ClaudeClient, error) {
	// sanity check
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable is not set")
	}

	// initialize client
	c := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
	)
	return &ClaudeClient{client: &c}, nil
}

func (c *ClaudeClient) Generate(ctx context.Context, prompt string) (string, error) {
	message, err := c.client.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeSonnet4_5,
		MaxTokens: int64(1024),
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	})

	if err != nil {
		return "", err
	}

	if len(message.Content) > 0 {
		return message.Content[0].Text, nil
	}

	return "", fmt.Errorf("no choices returned from Claude")
}
