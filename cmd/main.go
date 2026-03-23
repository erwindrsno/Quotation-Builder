package main

import (
	"github.com/erwindrsno/Quotation-Builder/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Company struct {
	id   int
	name string
}

type User struct {
	id       uuid.UUID
	username string
	password string
	name     string
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users/:name", user.MiddlewareOne(), user.Read)
	router.Run() // listens on 0.0.0.0:8080 by default
}
