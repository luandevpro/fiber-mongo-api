package models

import "time"

type User struct {
	ID        uint64    `json:"id",gorm:"primaryKey"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Password  string    `json:"password"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"crearedAt",sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt time.Time `json:"updatedAt",sql:"DEFAULT:'current_timestamp'"`
}
