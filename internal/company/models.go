package company

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type CreateReq struct {
	Name string `json:"name" binding:"required"`
}

type ReadReq struct {
	Name    string `form:"name,omitempty"`
	Compact bool   `form:"compact,omitempty"`
	Page    int    `form:"page,omitempty,default=1"`
	Size    int    `form:"size,omitempty,default=10"`
}
