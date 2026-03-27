package main

import (
	"log"

	"github.com/erwindrsno/Quotation-Builder/internal/company"
	"github.com/erwindrsno/Quotation-Builder/internal/database"
	"github.com/erwindrsno/Quotation-Builder/internal/role"
	"github.com/erwindrsno/Quotation-Builder/internal/user"
	"github.com/erwindrsno/Quotation-Builder/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.InitDB()
	defer db.Close()
	var hasher util.Hasher = util.ArgonHasher{}
	userRepo := &user.Repository{DB: db}
	userSvc := &user.Service{Repo: userRepo, Hasher: hasher}
	userCtrl := &user.Controller{Svc: userSvc}

	roleRepo := &role.Repository{DB: db}
	roleSvc := &role.Service{Repo: roleRepo}
	roleCtrl := &role.Controller{Svc: roleSvc}

	companyRepo := &company.Repository{DB: db}
	companySvc := &company.Service{Repo: companyRepo}
	companyCtrl := &company.Controller{Svc: companySvc}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.GET("/users/:name", user.MiddlewareOne(), user.Read)
	router.GET("/users", userCtrl.Read)
	router.POST("/users", userCtrl.Create)
	router.POST("/users/login", userCtrl.Login)
	router.GET("/roles", roleCtrl.Read)
	router.POST("/roles", roleCtrl.Create)
	router.GET("/companies", companyCtrl.Read)
	router.Run() // listens on 0.0.0.0:8080 by default
}
