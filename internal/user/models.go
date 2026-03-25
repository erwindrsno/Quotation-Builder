package user

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password,omitempty"`
	Name      string     `json:"name"`
	RoleId    int        `json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// DTOs
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

type Read struct {
	// Name filter also represents username
	Name string `form:"name"`
	Page int    `form:"page,default=1"`
	Size int    `form:"size,default=10"`
}
