package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"go-chat/router"
)

// wxfd7246ec04b25911 appID
//  68dafe32994254e53b0a320e56c749b0 as

func InitRouter() {
	engine := gin.Default()
	engine.Use(cors.Default())
	v1 := engine.Group("/v1")
	router.QuestionRouter(v1)

	engine.Run("0.0.0.0:8081")
}
