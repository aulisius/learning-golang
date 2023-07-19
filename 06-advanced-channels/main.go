package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Set up the OpenAI API client
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Create a channel to send the search query
	queries := make(chan string)
	completions := make(chan string)

	// Start the search function in a separate goroutine
	go search(client, queries, completions)

	fmt.Print("User: ")
	var query string
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v", err)
	}
	// Send the search query to the channel
	queries <- query

	for completion := range completions {
		fmt.Print("ChatGPT: ")
		fmt.Println(completion)
		// Get the search query from the user
		var query string
		fmt.Print("User: ")
		reader := bufio.NewReader(os.Stdin)
		query, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read input: %v", err)
		}
		// Send the search query to the channel
		queries <- query
	}
}

func search(client *openai.Client, queries chan string, completions chan string) {
	for query := range queries {
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: query,
					},
				},
			},
		)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}
		completions <- resp.Choices[0].Message.Content
	}
}
