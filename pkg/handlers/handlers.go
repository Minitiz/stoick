package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handlers struct{}

var Gin = new(Handlers)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("requestID", requestID)
		c.Next()
	}
}

// HandlerGin ...
func (h *Handlers) HandlerGin(r *gin.Engine) {
	r.Use(RequestIDMiddleware())
	h.urlShortener(r)
}
