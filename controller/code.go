package controller

//定义业务的状态码

type ServiceCode int64

const (
	CodeSuccess ServiceCode = 100
	CodeBusy    ServiceCode = 104
	//...
)

var msgFlags = map[ServiceCode]string{
	CodeSuccess: "成功",
	CodeBusy:    "业务繁忙，请稍后重试",
	//...
}

func (c ServiceCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeBusy]
}
