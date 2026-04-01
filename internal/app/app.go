package app

import (
	"database/sql"

	"github.com/erwindrsno/Quotation-Builder/internal/client"
	"github.com/erwindrsno/Quotation-Builder/internal/company"
	"github.com/erwindrsno/Quotation-Builder/internal/database"
	"github.com/erwindrsno/Quotation-Builder/internal/item_status"
	"github.com/erwindrsno/Quotation-Builder/internal/product"
	"github.com/erwindrsno/Quotation-Builder/internal/role"
	"github.com/erwindrsno/Quotation-Builder/internal/user"
	"github.com/erwindrsno/Quotation-Builder/internal/util"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func New() *App {
	db := database.InitDB()
	router := gin.Default()

	a := &App{
		DB:     db,
		Router: router,
	}

	a.setRoutes()
	return a
}

func (a *App) setRoutes() {
	a.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	var hasher util.Hasher = util.ArgonHasher{}

	userCtrl := user.New(a.DB, hasher)
	roleCtrl := role.New(a.DB)
	companyCtrl := company.New(a.DB)
	clientCtrl := client.New(a.DB)
	itemStatusCtrl := item_status.New(a.DB)
	productCtrl := product.New(a.DB)

	api := a.Router.Group("/api")
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
}

func (a *App) Run(port string) {
	a.Router.Run(port)
}

func (a *App) CloseDB() {
	a.DB.Close()
}
