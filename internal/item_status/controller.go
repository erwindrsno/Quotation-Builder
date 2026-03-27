package item_status

import (
	"net/http"

	"github.com/erwindrsno/Quotation-Builder/internal/responses"
	"github.com/gin-gonic/gin"
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
	var req ReadReq
	if itemStatuses, err := ctrl.Svc.ReadList(c.Request.Context(), &req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
	} else {
		responses.Success(c, http.StatusOK, gin.H{"item_statuses": itemStatuses})
	}
	return
}
