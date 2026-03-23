package user

import (
	"fmt"
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
	fmt.Printf("read.damn\n")
	name := ctx.Param("name")
	ctx.JSON(200, gin.H{"message": name})
}

func MiddlewareOne() gin.HandlerFunc {
	fmt.Printf("thisis middlewareOne, not in func\n")
	return func(ctx *gin.Context) {
		fmt.Printf("This is middlewareOne, in func\n")
		ctx.Next()
		fmt.Printf("m1 - after handler\n")
	}
}
