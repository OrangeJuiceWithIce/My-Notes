package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenExpireDuration = time.Second * 10 // 10秒过期
var CustomSecret = []byte("your-secret")     //JWT签名的密钥
// JWT存储的结构体
type CustomClaims struct {
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        //内嵌了一个标准的claim，大概包含了过期时间之类的信息
}

func GenToken(UserID int64, Username string) (string, error) {
	claims := CustomClaims{
		UserID,
		Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "lh",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	var claims = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	//检验claims对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
