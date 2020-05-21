//File  : UserDto.go
//Author: duanhaobin
//Date  : 2020/5/21

package dto

import "ginEssential-hb/model"

type UserDto struct {
	// dto 属性记得首字母大写
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}

}
