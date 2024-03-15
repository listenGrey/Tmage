package status

//定义业务的状态码

type Code int64

const (
	StatusSuccess Code = 100
	StatusBusy    Code = 104
	//...
	FuncCheckExistence Code = 700
)

var msgFlags = map[Code]string{
	StatusSuccess: "成功",
	StatusBusy:    "业务繁忙，请稍后重试",
	//...
	FuncCheckExistence: "校验该用户是否已经注册过",
}

func (c Code) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[StatusBusy]
}
