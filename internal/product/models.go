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
