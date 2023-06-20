package main

import (
	"log"

	"github.com/chennakt9/api-gateway.git/pkg/auth"
	"github.com/chennakt9/api-gateway.git/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}