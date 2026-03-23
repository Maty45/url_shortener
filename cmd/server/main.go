package main

import (
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewURLRepository()
	service := service.NewURLService(repo)
	handler := handler.NewURLHandler(service)

	r.POST("/shorten", handler.Create)
	r.GET("/:code", handler.Redirect)

	r.Run(":8080")

}
