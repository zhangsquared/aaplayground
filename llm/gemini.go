package llm

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const GeminiModel = "gemini-2.5-flash"

type GeminiClient struct {
	client *genai.Client
}

func InitGeminiClient() (*GeminiClient, error) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &GeminiClient{client: client}, nil
}

func (g *GeminiClient) Generate(ctx context.Context, prompt string) (string, error) {
	result, err := g.client.Models.GenerateContent(
		ctx,
		GeminiModel,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	if len(result.Candidates) > 0 {
		return result.Text(), nil
	}

	return "", fmt.Errorf("no candidates returned from Gemini")
}
