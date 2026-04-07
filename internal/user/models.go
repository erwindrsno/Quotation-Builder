package user

import (
	"time"

	"github.com/erwindrsno/Quotation-Builder/internal/role"
)

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password,omitempty"`
	Name      string     `json:"name"`
	RoleId    int        `json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Role      *role.Role `json:"role"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) ToReadResponse() ReadResponse {
	resp := ReadResponse{
		Id:        u.Id,
		Username:  u.Username,
		Name:      u.Name,
		RoleName:  u.Role.Name,
		CreatedAt: u.CreatedAt,
	}

	if u.Role != nil {
		resp.RoleName = u.Role.Name
	} else {
		resp.RoleName = "N/A"
	}

	if u.UpdatedAt != nil {
		resp.UpdatedAt = *u.UpdatedAt
	}
	return resp
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

type ReadResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
