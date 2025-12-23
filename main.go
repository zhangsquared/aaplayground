package main

import (
	"context"
	"fmt"
	"strings"

	"aa.zhangsquared.com/llm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// 1. Define Flags using pflag
	pflag.String("model", "gemini", "LLM model to use (gemini, chatgpt, claude)")
	pflag.Parse()

	// 2. Bind pflags to Viper
	// This tells Viper to look at the command-line flags
	viper.BindPFlags(pflag.CommandLine)

	// 3. Setup Environment Variables
	viper.SetEnvPrefix("APP") // Env vars look like APP_MODEL
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// 4. Setup Config File (Optional)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig() // Ignore error if file doesn't exist
}

func main() {
	// Viper handles the precedence: Flag > Env > Config > Default
	llmType := viper.GetString("model")

	var provider llm.ProviderInterface
	var err error

	switch llmType {
	case "gemini":
		provider, err = llm.InitGeminiClient()
	case "chatgpt":
		provider, err = llm.InitChatGPTClient()
	case "claude":
		provider, err = llm.InitClaudeClient()
	default:
		fmt.Printf("Unsupported LLM: %s\n", llmType)
		return
	}

	if err != nil {
		fmt.Printf("Init error: %v\n", err)
		return
	}

	resp, err := provider.Generate(context.Background(), "Hello World!")
	if err != nil {
		fmt.Printf("Generate error: %v\n", err)
		return
	}
	fmt.Println(resp)
}
