package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/util"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func encryptPwd(pwdByte []byte) (res string) {
	hashedPassword := md5.Sum(pwdByte)
	return hex.EncodeToString(hashedPassword[:])
}

func Register(client *models.RegisterFrom) (err error) {
	//call is grpc client
	// 1. Judging client exist
	// grpc server return a existence flag
	//infoCode

	existence := util.CheckExistence(client.Email)
	if existence == status.StatusSuccess {
		return errors.New(status.StatusSuccess.Msg())
	} else if existence == status.StatusBusy {
		return errors.New(status.StatusBusy.Msg())
	}

	// 2. generate ID and encrypt password
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	userId := node.Generate()
	pwdByte := []byte(client.Password)
	userPwd := encryptPwd(pwdByte)

	//create a user
	user := &models.User{
		UserID:   userId.Int64(),
		Email:    client.Email,
		UserName: client.UserName,
		Password: userPwd,
	}

	// 3. call a grpc client, send user info

	return fmt.Errorf("register success, ID: %d,email: %s,username: %s,password: %s", user.UserID, user.Email, user.UserName, user.Password)
}

/*
func Login(form *models.LoginForm) (atoken, rtoken string, err error) {
	user := &models.User{
		Email:    form.Email,
		Password: form.Password,
	}

	// call a grpc client, send user info
	if err := err != nil {

	}

	// 生成JWT
	return "", "", err
}
*/
