//File  : jwt.go
//Author: duanhaobin
//Date  : 2020/5/20

package common

import (
	"ginEssential-hb/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("xiang_shou_hb")

type Claims struct {
	UserID int
	jwt.StandardClaims
}

// 注册成功后发放 token
func ReleaseToken(user model.User) (string, error) {
	// 设置过期时间  15 天
	expirationTime := time.Now().Add(7 * 24 * 15)

	claims := &Claims{
		UserID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 生效时间，立即生效
			Issuer:    "xiangshou",           // 发放机构
			Subject:   "user token",          // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeString, err := token.SignedString(jwtKey) // 使用秘钥生成token
	if err != nil {
		return "", err
	}
	return tokeString, nil
}
