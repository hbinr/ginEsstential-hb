//File  : UserController.go
//Author: duanhaobin
//Date  : 2020/5/20

package controller

import (
	"ginEssential-hb/common"
	"ginEssential-hb/model"
	"ginEssential-hb/util"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Register:注册
func Register(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
		return
	}

	if len(name) == 0 {
		// 如果用户未输入用户名，自定义用户名
		name = util.RandomString(10)
	}
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已注册"})
		return
	}
	// 密码加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
	}
	// 创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPwd),
	}
	DB.Create(&user)
	// 返回响应
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

// Login:登录
func Login(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")
	// 验证数据
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
		return
	}

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	// 验证手机号是否存在
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("生成token异常，err:%v", err)
		return
	}

	// 响应
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": token, "msg": "登录成功"})
}

// Info:用户信息
func Info(c *gin.Context) {
	// 正确的处理:用户信息必须是经过认证的，因此从上下文中获取用户信息
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})

}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
