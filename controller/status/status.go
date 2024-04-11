package status

//定义业务的状态码

type Code int64

const (
	// 可以暴露给外部的错误码
	StatusSuccess           Code = 1000
	StatusInvalidParams     Code = 1001
	StatusUserExist         Code = 1002
	StatusUserNotExist      Code = 1003
	StatusInvalidPwd        Code = 1004
	StatusBusy              Code = 1005
	StatusInvalidToken      Code = 1006
	StatusInvalidAuthFormat Code = 1007
	StatusNotLogin          Code = 1008
	StatusShareInvalid      Code = 1009

	// 内部错误码
	StatusInvalidGenID       Code = 1100
	StatusRegisterERR        Code = 1101
	StatusConnGrpcServerERR  Code = 1102
	StatusRecvGrpcSerInfoERR Code = 1103
	StatusConnDBERR          Code = 1104
	StatusImagesNotFound     Code = 1105
	StatusNoTag              Code = 1106
	StatusKafkaSendERR       Code = 1107
	StatusKafkaReceiveERR    Code = 1108
)

var msgFlags = map[Code]string{
	StatusSuccess:           "成功",
	StatusInvalidParams:     "请求参数错误",
	StatusUserExist:         "用户已存在",
	StatusUserNotExist:      "用户不存在",
	StatusInvalidPwd:        "用户名或密码错误",
	StatusBusy:              "业务繁忙，请稍后重试",
	StatusInvalidToken:      "无效的Token",
	StatusInvalidAuthFormat: "认证格式有误",
	StatusNotLogin:          "未登录",
	StatusShareInvalid:      "无效的分享链接",

	StatusInvalidGenID:       "生成ID失败",
	StatusRegisterERR:        "用户注册失败",
	StatusConnGrpcServerERR:  "无法连接到gRpc服务器",
	StatusRecvGrpcSerInfoERR: "从gRpc服务器获取信息失败",
	StatusConnDBERR:          "无法连接到数据库",
	StatusImagesNotFound:     "没有找到图片",
	StatusNoTag:              "没有标签",
	StatusKafkaSendERR:       "向kafka中发送数据失败",
	StatusKafkaReceiveERR:    "从kafka中获取数据失败",
}

func (c Code) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[StatusBusy]
}

func (c Code) Code() int64 {
	return int64(c)
}
