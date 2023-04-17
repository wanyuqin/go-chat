package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type Question struct {
	Content string `json:"content"`
}

func PostQuestion(c *gin.Context) {
	q := Question{}
	err := c.ShouldBind(&q)
	if err != nil {
		c.JSON(200, err)
		c.Abort()
		return

	}
	fmt.Println(q.Content)
	answer := chat(q.Content)
	resp := make(map[string]interface{})
	resp["data"] = answer
	c.JSON(200, answer)
}

func chat(content string) string {
	key := os.Getenv("CHAT_KEY")

	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}
