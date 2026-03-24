package user

import (
	"net/http"

	"github.com/erwindrsno/Quotation-Builder/internal/responses"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Svc *Service
}

func (ctrl *Controller) Create(c *gin.Context) {
	var json Register

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if errCreate := ctrl.Svc.Create(c.Request.Context(), &json); errCreate != nil {
		responses.Fail(c, http.StatusBadRequest, errCreate.Error())
	} else {
		responses.Success(c, http.StatusCreated, gin.H{"success": true})
	}
}

func (ctrl *Controller) Read(c *gin.Context) {
	// var json Read
}
