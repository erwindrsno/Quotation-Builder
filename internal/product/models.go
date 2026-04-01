package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id                uuid.UUID  `json:"id"`
	Name              string     `json:"name"`
	PartNumber        string     `json:"part_number"`
	Description       string     `json:"description"`
	BaseUnit          string     `json:"base_unit"`
	ManufacturerBrand string     `json:"manufacturer_brand"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

type CreateReq struct {
	Name              string `json:"name" binding:"required"`
	PartNumber        string `json:"part_number" binding:"required"`
	Description       string `json:"description" binding:"required"`
	BaseUnit          string `json:"base_unit" binding:"required"`
	ManufacturerBrand string `json:"manufacturer_brand" binding:"required"`
}

type ReadReq struct {
	Name    string `form:"name,omitempty"`
	Compact bool   `form:"compact,omitempty"`
	Page    int    `form:"page,omitempty,default=1"`
	Size    int    `form:"size,omitempty,default=10"`
}
