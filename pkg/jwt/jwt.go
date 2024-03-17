package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	STD    jwt.StandardClaims
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}

// secret
var secret = []byte("你知道我想说什么")

func keyGen(_ *jwt.Token) (key interface{}, err error) {
	return secret, nil
}

// 定义JWT的过期时间
const ShortTokenExpireDuration = time.Minute * 30
const LongTokenExpireDuration = time.Hour * 24 * 15

func GenToken(userID int64) (aToken, rToken string, err error) {
	claim := Claims{
		UserID: userID,
		STD: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ShortTokenExpireDuration).Unix(),
			Issuer:    "listenGrey",
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secret)

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(LongTokenExpireDuration).Unix(),
		Issuer:    "listenGrey",
	}).SignedString(secret)
	return
}

// 解析token
func ParseToken(oriToken string) (claim *Claims, err error) {
	var token *jwt.Token
	claim = new(Claims)
	token, err = jwt.ParseWithClaims(oriToken, claim, keyGen)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

// 刷新aToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	if _, err = jwt.Parse(rToken, keyGen); err != nil {
		return
	}

	var claims Claims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyGen)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
