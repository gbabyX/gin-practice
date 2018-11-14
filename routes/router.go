package routes

import "github.com/gin-gonic/gin"

func init() {

}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/ping", controllers.Pong())
	return router
}
