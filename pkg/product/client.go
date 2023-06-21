package product

import (
	"fmt"

	"github.com/chennakt9/api-gateway.git/pkg/config"
	"github.com/chennakt9/api-gateway.git/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect..", err)
	}

	return pb.NewProductServiceClient(cc)
}