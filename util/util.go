package util

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/pkg/grpc"
	"fmt"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"

	"context"
)

func CheckExistence(email string) status.Code {
	client := grpc.ClientServer(grpc.CheckExistence)
	if client == status.StatusConnGrpcServerErr {
		return status.StatusConnGrpcServerErr
	}
	sendEmail := &userInfo.RegisterEmail{Email: email}
	res, err := client.(userInfo.CheckExistenceClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		fmt.Printf("Failed to receive info from gRpc server; %v\n", err)
		return status.StatusRecvGrpcSerInfoErr
	}
	if exist := res.Exsit; exist {
		return status.StatusUserExist
	} else {
		return status.StatusSuccess
	}
}

func Register(user *models.User) status.Code {
	client := grpc.ClientServer(grpc.Register)
	if client == status.StatusConnGrpcServerErr {
		return status.StatusConnGrpcServerErr
	}
	sendUser := &userInfo.RegisterForm{
		UserID:   user.UserID,
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}
	res, err := client.(userInfo.RegisterInfoClient).Register(context.Background(), sendUser)
	if err != nil {
		fmt.Printf("Failed to receive info from gRpc server; %v\n", err)
		return status.StatusRecvGrpcSerInfoErr
	}
	if sta := res.Success; sta {
		return status.StatusSuccess
	} else {
		return status.StatusRegisterErr
	}
}

func LoginCheck(user *models.User) (code status.Code, userID int64) {
	client := grpc.ClientServer(grpc.LoginCheck)
	if client == status.StatusConnGrpcServerErr {
		return status.StatusConnGrpcServerErr, 0
	}
	sendUser := &userInfo.LoginForm{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := client.(userInfo.LoginCheckClient).LoginCheck(context.Background(), sendUser)
	if err != nil {
		fmt.Printf("Failed to receive info from gRpc server; %v\n", err)
		return status.StatusRecvGrpcSerInfoErr, 0
	}
	sta := res.Info
	userID = res.UserID
	if sta == int64(status.StatusSuccess) {
		return status.StatusSuccess, userID
	} else if sta == int64(status.StatusUserNotExist) {
		return status.StatusUserNotExist, 0
	} else if sta == int64(status.StatusInvalidPwd) {
		return status.StatusInvalidPwd, 0
	} else {
		return status.StatusBusy, 0
	}
}
