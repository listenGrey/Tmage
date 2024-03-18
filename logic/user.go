package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/pkg/jwt"
	"Tmage/util"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/bwmarrin/snowflake"
)

func encryptPwd(pwdByte []byte) (res string) {
	hashedPassword := md5.Sum(pwdByte)
	return hex.EncodeToString(hashedPassword[:])
}

func Register(client *models.RegisterFrom) (err error) {
	// 1. 将注册邮箱通过gRpc发送到gRpc服务器去判断用户是否存在
	existence := util.CheckExistence(client.Email)
	if existence != status.StatusSuccess {
		return errors.New(existence.Msg())
	}

	// 2. 生成ID，对密码加密
	node, err := snowflake.NewNode(1)
	if err != nil {
		return errors.New(status.StatusInvalidGenID.Msg())
	}
	userId := node.Generate()
	pwdByte := []byte(client.Password)
	userPwd := encryptPwd(pwdByte)

	// 创建一个用户
	user := &models.User{
		UserID:   userId.Int64(),
		Email:    client.Email,
		UserName: client.UserName,
		Password: userPwd,
	}

	// 3. call a grpc client, send user info
	res := util.Register(user)
	if res != status.StatusSuccess {
		return errors.New(res.Msg())
	}
	return nil
}

func Login(form *models.LoginForm) (user *models.User, err error) {
	// 对密码加密
	pwdByte := []byte(form.Password)
	userPwd := encryptPwd(pwdByte)

	user = &models.User{
		Email:    form.Email,
		Password: userPwd,
	}

	// 将登录信息通过gRpc发送到gRpc服务器去判断用户和密码是否正确
	info, userID := util.LoginCheck(user)
	if info != status.StatusSuccess {
		return nil, errors.New(info.Msg())
	}
	user.UserID = userID

	// 生成JWT
	aToken, rToken, err := jwt.GenToken(user.UserID)
	if err != nil {
		return
	}
	user.AccessToken = aToken
	user.RefreshToken = rToken
	return
}
