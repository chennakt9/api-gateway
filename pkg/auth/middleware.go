package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/chennakt9/api-gateway.git/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("you are not authorized"))
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("you are not authorized"))
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("you are not authorized"))
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}