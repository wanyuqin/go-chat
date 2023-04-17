package initialize

import (
	"github.com/gin-gonic/gin"

	"go-chat/handler"
)

// wxfd7246ec04b25911 appID
//  68dafe32994254e53b0a320e56c749b0 as

func InitRouter() {
	engine := gin.Default()

	engine.POST("/question", handler.PostQuestion)

	engine.Run("0.0.0.0:8081")
}
