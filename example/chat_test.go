package example

import (
	"context"
	"fmt"
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

func TestChat(t *testing.T) {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "chatGPT 默认是会给账号分配18美金吗，是如何计费的",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
