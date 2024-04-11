package util

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/pkg/grpc"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"

	"context"
)

func CheckExistence(email string) status.Code {
	client := grpc.UserClientServer(grpc.CheckExistence)
	if client == status.StatusConnGrpcServerERR {
		return status.StatusConnGrpcServerERR
	}
	sendEmail := &userInfo.RegisterEmail{Email: email}
	res, err := client.(userInfo.CheckExistenceClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		return status.StatusRecvGrpcSerInfoERR
	}
	exist := res.Exist
	info := res.Info

	if info == status.StatusConnDBERR.Code() {
		return status.StatusConnDBERR
	} else if info == status.StatusBusy.Code() {
		return status.StatusBusy
	}

	if exist {
		return status.StatusUserExist
	}

	return status.StatusSuccess
}

func Register(user *models.User) status.Code {
	client := grpc.UserClientServer(grpc.Register)
	if client == status.StatusConnGrpcServerERR {
		return status.StatusConnGrpcServerERR
	}
	sendUser := &userInfo.RegisterForm{
		UserID:   user.UserID,
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}
	res, err := client.(userInfo.RegisterInfoClient).Register(context.Background(), sendUser)
	if err != nil {
		return status.StatusRecvGrpcSerInfoERR
	}

	sta := res.Success
	info := res.Info

	if info == status.StatusConnDBERR.Code() {
		return status.StatusConnDBERR
	} else if info == status.StatusBusy.Code() {
		return status.StatusBusy
	}

	if !sta {
		return status.StatusRegisterERR
	}

	return status.StatusSuccess
}

func LoginCheck(user *models.User) (code status.Code, userID int64) {
	client := grpc.UserClientServer(grpc.LoginCheck)
	if client == status.StatusConnGrpcServerERR {
		return status.StatusConnGrpcServerERR, 0
	}
	sendUser := &userInfo.LoginForm{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := client.(userInfo.LoginCheckClient).LoginCheck(context.Background(), sendUser)
	if err != nil {
		return status.StatusRecvGrpcSerInfoERR, 0
	}
	sta := res.Info
	userID = res.UserID
	if sta == status.StatusConnDBERR.Code() {
		return status.StatusConnDBERR, 0
	} else if sta == status.StatusUserNotExist.Code() {
		return status.StatusUserNotExist, 0
	} else if sta == status.StatusInvalidPwd.Code() {
		return status.StatusInvalidPwd, 0
	} else if sta == status.StatusBusy.Code() {
		return status.StatusBusy, 0
	}

	return status.StatusSuccess, userID
}
