package repository

import "gorm.io/gorm"

type Admin struct {
	db *gorm.DB
}

func NewAdmin(dbCrud *gorm.DB) Admin {
	return Admin{
		db: dbCrud,
	}
}

type AdminInterfaceRepo interface {
}
