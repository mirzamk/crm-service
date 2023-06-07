package repository

import "gorm.io/gorm"

type ActorRepository struct {
	db *gorm.DB
}

func ActorNewRepo(db *gorm.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

type ActorInterfaceRepo interface {
}
