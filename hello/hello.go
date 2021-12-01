package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello() *gin.Engine{
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	return r
}
