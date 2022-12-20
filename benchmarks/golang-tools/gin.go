package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang_benchmarks/handlers"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	h := gin.New()
	h.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})
	h.POST("/", func(c *gin.Context) {
		var req Req
		if err := c.Bind(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if len(req.Data) == 0 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.String(http.StatusOK, req.Data[len(req.Data)-1])
	})
	h.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})
	handlers.RegisterHandler("gin", h)
}
