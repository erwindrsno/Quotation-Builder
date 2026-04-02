package quotation

import (
	"time"

	"github.com/google/uuid"
)

type Quotation struct {
	Id                  uuid.UUID  `json:"id"`
	Subject             string     `json:"subject"`
	QuoteNumber         string     `json:"quote_number"`
	Validity            time.Time  `json:"validity"`
	DeliveryTime        *time.Time `json:"delivery_time"`
	Deadline            *time.Time `json:"deadline"`
	DeliveryDestination string     `json:"delivery_destination"`
	TermsOfPayment      string     `json:"terms_of_payment"`
	Notes               string     `json:"notes"`
	Discount            float32    `json:"discount"`
	ClientId            uuid.UUID  `json:"client_id"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
}

type CreateReq struct {
	Subject             string     `json:"subject" binding:"required"`
	QuoteNumber         string     `json:"quote_number" binding:"required"`
	Validity            time.Time  `json:"validity" binding:"required"`
	DeliveryTime        *time.Time `json:"delivery_time"`
	Deadline            *time.Time `json:"deadline"`
	DeliveryDestination string     `json:"delivery_destination" binding:"required"`
	TermsOfPayment      string     `json:"terms_of_payment" binding:"required"`
	Notes               string     `json:"notes" binding:"required"`
	Discount            float32    `json:"discount" binding:"required"`
	ClientId            uuid.UUID  `json:"client_id" binding:"required"`
}

type ReadAllReq struct {
	Search string `form:"search,omitempty"`
	Page   int    `form:"page,omitempty,default=1"`
	Size   int    `form:"size,omitempty,default=10"`
}

type ReadOneReq struct {
	Id uuid.UUID `form:"id"`
}
