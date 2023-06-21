package main

import (
	"log"

	"github.com/chennakt9/api-gateway.git/pkg/auth"
	"github.com/chennakt9/api-gateway.git/pkg/config"
	"github.com/chennakt9/api-gateway.git/pkg/order"
	"github.com/chennakt9/api-gateway.git/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}