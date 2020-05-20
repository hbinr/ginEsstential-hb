//File  : UserController.go
//Author: duanhaobin
//Date  : 2020/5/20

package controller

import (
	"ginEssential-hb/common"
	"ginEssential-hb/model"
	"ginEssential-hb/util"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// 注册
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

	// 创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&user)
	// 返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
