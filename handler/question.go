package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func ChatStream(conn *websocket.Conn, content string) {
	key := os.Getenv("CHAT_KEY")

	c := openai.NewClient(key)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)

	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Finished"))
			return
		}
		// 发送消息
		err = conn.WriteMessage(websocket.TextMessage, []byte(response.Choices[0].Delta.Content))
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

func Conversation(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := conn.ReadMessage()
		fmt.Println(string(message))
		if err != nil {
			log.Println(err)
			return
		}
		ChatStream(conn, string(message))
	}

}
