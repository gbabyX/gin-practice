package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
