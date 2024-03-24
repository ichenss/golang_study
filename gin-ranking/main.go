package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/url", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello gin")
	})

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
