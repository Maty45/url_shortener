package handler

import (
	"net/http"

	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

func (h *URLHandler) Create(c *gin.Context) {
	var req CreateURLRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	url := h.service.Create(req.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + url.Code,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, exists := h.service.GetByCode(code)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, url.Original)
}
