package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	username string
	password string
	name     string
}

// func create() {
//
// }

func Read(ctx *gin.Context) {

}
