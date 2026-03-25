package user

import (
	"log"
	"net/http"

	"github.com/erwindrsno/Quotation-Builder/internal/responses"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Svc *Service
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req Register

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if errCreate := ctrl.Svc.Create(c.Request.Context(), &req); errCreate != nil {
		responses.Fail(c, http.StatusBadRequest, errCreate.Error())
	} else {
		responses.Success(c, http.StatusCreated, gin.H{"success": true})
	}
	return
}

func (ctrl *Controller) Read(c *gin.Context) {
	var req Read

	if err := c.ShouldBindQuery(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("name=%s, page=%d, size=%d", req.Name, req.Page, req.Size)

	if users, errRead := ctrl.Svc.Read(c.Request.Context(), &req); errRead != nil {
		responses.Fail(c, http.StatusBadRequest, errRead.Error())
	} else {
		responses.Success(c, http.StatusOK, gin.H{"users": users})
	}
	return
}
