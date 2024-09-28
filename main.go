package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/predictionguard/go-client"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	host := "https://api.predictionguard.com"
	apiKey := os.Getenv("PGKEY")

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	maxTokens := 1000
	temperature := float32(0.3)

	input := client.ChatSSEInput{
		Model: "Hermes-2-Pro-Llama-3-8B", // Updated model name
		Messages: []client.ChatInputMessage{
			{
				Role:    client.Roles.System,
				Content: "You are a helpful coding assistant.",
			},
			{
				Role:    client.Roles.User,
				Content: "Write a Go program that prints out random numbers.",
			},
		},
		MaxTokens:   &maxTokens,
		Temperature: &temperature,
	}

	ch := make(chan client.ChatSSE, 100)

	err := cln.ChatSSE(ctx, input, ch)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	for resp := range ch {
		for _, choice := range resp.Choices {
			fmt.Print(choice.Delta.Content)
		}
	}

	return nil
}
