package routes

import (
	"../service/index"
	"github.com/gin-gonic/gin"
)

func init() {

}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/ping", index.Pong)
	return router
}
