//File  : AuthMiddleware.go
//Author: duanhaobin
//Date  : 2020/5/21

// 用户登录认证中间件

package middleware

import (
	"ginEssential-hb/common"
	"ginEssential-hb/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 基于Gin的中间件,返回函数是一个 HandleFunc
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header中获取 authorization
		tokenStr := c.GetHeader("Authorization")

		// 验证格式：非空且以 "Bearer "  有空格，7个字符
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() // 不再执行中间件之后的函数
			return
		}

		// toke 有效，提取header中token的有效部分
		tokenStr = tokenStr[7:]
		// 解析 token
		token, claims, err := common.ParseToken(tokenStr)
		// 解析失败或token无效，也返回权限不足
		if err != nil || !token.Valid {
			log.Println("token解析失败或无效，err:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() // 不再执行中间件之后的函数
			return
		}

		// token 通过验证，获取用户中的userId
		userId := claims.UserID

		// 验证用户是否存在，用户不存在(或被消除 )，这个token无效
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() // 不再执行中间件之后的函数
			return
		}

		// 用户存在，将用户的信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
