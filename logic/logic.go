package logic

import (
	"Tmage/models"
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
	//call is gRpc client
	// 1. Judging client exist
	// gRpc server return a existence flag
	tempUser := models.RegisterFrom{
		UserName:   "",
		Email:      "",
		Password:   "",
		RePassword: "",
	}
	exist := false
	if client.UserName == tempUser.UserName && client.Email == tempUser.Email && client.Password == tempUser.Password && client.RePassword == tempUser.RePassword {
		exist = true
	}
	if exist {
		return errors.New("用户已存在")
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

	return fmt.Errorf("register success, ID: %d,email: %s,username: %s,password: %s", user.UserID, user.Email, user.UserName, user.Password)

	// 3. call a gRpc client, send user info
}

/*
func Login(form *models.LoginForm) (atoken, rtoken string, err error) {
	user := &models.User{
		Email:    form.Email,
		Password: form.Password,
	}

	// call a gRpc client, send user info
	if err := err != nil {

	}

	// 生成JWT
	return "", "", err
}
*/
