package company

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReq struct {
	Name string `json:"name" binding:"required"`
}

type ReadReq struct {
	Name string `json:"name" binding:"required"`
	Page int    `json:"page,omitempty"`
	Size int    `json:"size,omitempty"`
}
