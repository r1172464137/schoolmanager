package serializer

import "school_manager/model"

type User struct {
	Uid      uint   `json:"uid" form:"uid"`                  // 用户ID
	UserName string `json:"user_name" form:"user_name" example:"FanOne"` // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"`                  // 创建
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		Uid:      user.Uid,
		UserName: user.Username,
	}
}
