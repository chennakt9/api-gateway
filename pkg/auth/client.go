package auth

import (
	"fmt"

	"github.com/chennakt9/api-gateway.git/pkg/auth/pb"
	"github.com/chennakt9/api-gateway.git/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect", err)
	}

	return pb.NewAuthServiceClient(cc)
}