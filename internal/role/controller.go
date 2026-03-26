package role

import (
	"github.com/erwindrsno/Quotation-Builder/internal/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Svc *Service
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.Svc.Create(c, &req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
	} else {
		responses.Success(c, http.StatusCreated, gin.H{"message": "created."})
	}
	return
}

func (ctrl *Controller) Read(c *gin.Context) {
	if roles, err := ctrl.Svc.Read(c.Request.Context()); err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		responses.Success(c, http.StatusOK, gin.H{"roles": roles})
	}
	return
}
