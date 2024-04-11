package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/pkg/jwt"
	"Tmage/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/bwmarrin/snowflake"
	"time"
)

func encryptPwd(pwdByte []byte) (res string) {
	hashedPassword := md5.Sum(pwdByte)
	return hex.EncodeToString(hashedPassword[:])
}

func Register(client *models.RegisterFrom) status.Code {
	// 1. 将注册邮箱通过gRpc发送到gRpc服务器去判断用户是否存在
	existence := util.CheckExistence(client.Email)
	if existence != status.StatusSuccess {
		return existence
	}

	// 2. 生成ID，对密码加密
	node, err := snowflake.NewNode(1)
	if err != nil {
		return status.StatusInvalidGenID
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

	// 3. 发送用户信息
	res := util.Register(user)
	if res != status.StatusSuccess {
		return res
	}

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    userId.Int64(),
		Operation: "register",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    res.Msg(),
		Data:      user,
	}
	util.OperationLog(*operationInfo)

	return status.StatusSuccess
}

func Login(form *models.LoginForm) (user *models.User, code status.Code) {
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
		return nil, info
	}
	user.UserID = userID

	// 生成JWT
	aToken, rToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, status.StatusBusy
	}
	user.AccessToken = aToken
	user.RefreshToken = rToken

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "login",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    info.Msg(),
		Data:      user,
	}
	util.OperationLog(*operationInfo)

	return user, status.StatusSuccess
}
