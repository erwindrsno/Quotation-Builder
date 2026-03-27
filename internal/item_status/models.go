package item_status

import (
	"time"

	"github.com/google/uuid"
)

type ItemStatus struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReq struct {
	Name string `json:"name"`
}

type ReadReq struct {
	Name string `json:"name"`
}
