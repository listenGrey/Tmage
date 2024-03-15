package util

import (
	"Tmage/controller/status"
	"Tmage/pkg/grpc"
	"fmt"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"

	"context"
)

func CheckExistence(email string) status.Code {
	client := grpc.ClientServer(status.FuncCheckExistence)
	if client == nil {
		fmt.Printf("grpc server did not connect\n")
		return status.StatusBusy
	}
	sendEmail := &userInfo.RegisterEmail{Email: email}
	res, err := client.(userInfo.CheckExistenceClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		fmt.Printf("Failed to receive: %v\n", err)
		return status.StatusBusy
	}
	if exist := res.Exsit; exist {
		return status.StatusSuccess
	} else {
		return status.StatusBusy
	}
}
