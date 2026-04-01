package main

import (
	"log"

	"github.com/erwindrsno/Quotation-Builder/internal/client"
	"github.com/erwindrsno/Quotation-Builder/internal/company"
	"github.com/erwindrsno/Quotation-Builder/internal/database"
	"github.com/erwindrsno/Quotation-Builder/internal/item_status"
	"github.com/erwindrsno/Quotation-Builder/internal/product"
	"github.com/erwindrsno/Quotation-Builder/internal/role"
	"github.com/erwindrsno/Quotation-Builder/internal/user"
	"github.com/erwindrsno/Quotation-Builder/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//First of all, remember to clear tmp files before running the app. If not, configuring line 17 to the path is necessary
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

	itemStatusRepo := &item_status.Repository{DB: db}
	itemStatusSvc := &item_status.Service{Repo: itemStatusRepo}
	itemStatusCtrl := &item_status.Controller{Svc: itemStatusSvc}

	clientRepo := &client.Repository{DB: db}
	clientSvc := &client.Service{Repo: clientRepo}
	clientCtrl := &client.Controller{Svc: clientSvc}

	productRepo := &product.Repository{DB: db}
	productSvc := &product.Service{Repo: productRepo}
	productCtrl := &product.Controller{Svc: productSvc}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.POST("/login", userCtrl.Login)
				protected := users.Group("")
				protected.Use(util.AuthMiddleware())
				{
					protected.GET("", userCtrl.Read)
					protected.POST("", userCtrl.Create)
				}
			}

			roles := v1.Group("/roles")
			roles.Use(util.AuthMiddleware())
			{
				roles.GET("", roleCtrl.Read)
				roles.POST("", roleCtrl.Create)
			}

			companies := v1.Group("/companies")
			{
				companies.GET("", companyCtrl.Read)
				companies.POST("", companyCtrl.Create)
			}

			itemStatuses := v1.Group("/item-statuses")
			{
				itemStatuses.GET("", itemStatusCtrl.Read)
				itemStatuses.POST("", itemStatusCtrl.Create)
			}

			clients := v1.Group("/clients")
			{
				clients.GET("", clientCtrl.Read)
				clients.POST("", clientCtrl.Create)
			}

			products := v1.Group("/products")
			{
				products.GET("", productCtrl.Read)
				products.POST("", productCtrl.Create)
			}
		}
	}
	router.Run() // listens on 0.0.0.0:8080 by default
}
