package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minitiz/url-shortener/pkg/db"
	handlers "github.com/minitiz/url-shortener/pkg/handlers"
	"github.com/minitiz/url-shortener/pkg/logger"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func init() {
	if err := db.New("localhost", "5432", "pierreyvesliegeois", "mysecretpassword", "urlshortener").CreateClient(); err != nil {
		panic(err)
	}
}

func main() {

	// register handler
	r := gin.Default()
	handlers.Gin.HandlerGin(r)
	log := logger.Logger
	if err := r.Run(":8080"); err != nil {
		log.Error("Run failed %s", zap.Error(err))
		return
	}
}
