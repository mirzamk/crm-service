package entity

import "github.com/mirzamk/crm-service/constant"

type Actor struct {
	GormModel
	Username   string
	Password   string
	RoleId     uint
	IsVerified constant.BoolType
	IsActive   constant.BoolType
	Role       *Role
}

func (Actor) TableName() string {
	return "actor"
}
