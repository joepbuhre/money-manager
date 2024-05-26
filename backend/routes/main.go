package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMainRouter(rg *gin.RouterGroup) {
	public := rg.Group("/")

	public.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("It is all working"))
	})
}
