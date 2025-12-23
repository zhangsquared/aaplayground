package llm

import (
	"context"
	"fmt"
)

// Provider defines the common "language" all your LLMs must speak
type ProviderInterface interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

// Agent is your "Brain" that doesn't care which LLM it's using
type Agent struct {
	Client ProviderInterface
}

func (a *Agent) RunTask(ctx context.Context, task string) {
	fmt.Printf("Agent is thinking about: %s\n", task)
	response, _ := a.Client.Generate(ctx, task)
	fmt.Println("Result:", response)
}
