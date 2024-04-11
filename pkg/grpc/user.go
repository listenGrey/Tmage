package grpc

import (
	"Tmage/controller/status"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc"
)

// 定义gRpc客户端服务器的类型码

type UserClient int64

const (
	CheckExistence UserClient = 2000
	Register       UserClient = 2001
	LoginCheck     UserClient = 2002
)

func UserClientServer(funcCode UserClient) (client interface{}) {
	conn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	if err != nil {
		return status.StatusConnGrpcServerERR
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
