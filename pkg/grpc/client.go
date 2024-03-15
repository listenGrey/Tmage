package grpc

import (
	"fmt"

	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc"

	"Tmage/controller/status"
)

func ClientServer(funcCode status.Code) (client interface{}) {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc server did not connect; %v\n", err)
		return nil
	}
	switch funcCode {
	case status.FuncCheckExistence:
		client = userInfo.NewCheckExistenceClient(conn)
	default:
		client = nil
	}
	return client
}
