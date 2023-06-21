package order

import (
	"fmt"

	"github.com/chennakt9/api-gateway.git/pkg/config"
	"github.com/chennakt9/api-gateway.git/pkg/order/pb"
	"github.com/chennakt9/api-gateway.git/pkg/order/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect", err)
	}

	return pb.NewOrderServiceClient(cc)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}