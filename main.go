package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	handler := gin.New()
	handler.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin delete method")
	})
	s := &http.Server{
		Addr:         ":3000",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
