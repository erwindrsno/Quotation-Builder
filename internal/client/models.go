package client

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CompanyId uuid.UUID  `json:"company_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CreateReq struct {
	Name      string    `json:"name" binding:"required"`
	CompanyId uuid.UUID `json:"company_id" binding:"required"`
}

type ReadReq struct {
	Name    string `form:"name,omitempty"`
	Compact bool   `form:"compact,omitempty"`
	Page    int    `form:"page,omitempty,default=1"`
	Size    int    `form:"size,omitempty,default=10"`
}
