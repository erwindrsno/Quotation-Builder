package role

import "time"

type Role struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReq struct {
	Name string `json:"name" binding:"required,min=3"`
}
