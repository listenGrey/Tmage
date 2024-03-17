package grpc

import (
	"Tmage/controller/status"
	"fmt"

	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc"
)

// 定义gRpc客户端服务器的类型码

type Client int64

const (
	CheckExistence Client = 2000
	Register       Client = 2001
	LoginCheck     Client = 2002
)

func ClientServer(funcCode Client) (client interface{}) {
	conn, err := grpc.Dial("", grpc.WithInsecure()) //server IP
	if err != nil {
		fmt.Println("cannot connect grpc server")
		return status.StatusConnGrpcServerErr
	}
	switch funcCode {
	case CheckExistence:
		client = userInfo.NewCheckExistenceClient(conn)
	case Register:
		client = userInfo.NewRegisterInfoClient(conn)
	case LoginCheck:
		client = userInfo.NewLoginCheckClient(conn)
	default:
		client = nil
	}
	return client
}
