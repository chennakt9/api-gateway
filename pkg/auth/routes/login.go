package routes

import (
	"context"
	"net/http"

	"github.com/chennakt9/api-gateway.git/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	// fmt.Println("login payload", b)

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email: b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(http.StatusOK, &res)
}