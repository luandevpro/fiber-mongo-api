package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id",gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

type Post struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      uint64 `json:"userId"`
	User        User   `json:"user"`
}
