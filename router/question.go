package router

import (
	"github.com/gin-gonic/gin"

	"go-chat/handler"
)

func QuestionRouter(e *gin.RouterGroup) {
	e.POST("/question", handler.PostQuestion)
	e.Any("/conversation", handler.Conversation)
}
