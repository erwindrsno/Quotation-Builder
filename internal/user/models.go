package user

import (
	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	username string
	password string
	name     string
}

// DTOs
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

type Read struct {
}
