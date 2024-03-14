package jwt

type JWT struct {
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	//jwt.//jwt官方类型
}

/*
//secret
var secret = []byte("你知道我想说什么")

func KeyGen(_ *jwt.Token) (key interface{},err error) {
	return secret,nil
}

//定义JWT的过期时间
const TokenExpireDuration = time.Minute * 30

func GenToken(userID uint64,email string) (atoken,rtoken string,err error) {

}*/
