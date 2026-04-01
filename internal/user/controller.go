package user

import (
	"errors"
	"github.com/erwindrsno/Quotation-Builder/internal/responses"
	"github.com/gin-gonic/gin"
	"net/http"
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

	if err := ctrl.Svc.Create(c.Request.Context(), &req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
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

	if users, err := ctrl.Svc.Read(c.Request.Context(), &req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
	} else {
		responses.Success(c, http.StatusOK, gin.H{"users": users})
	}
	return
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req Login

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	if ss, err := ctrl.Svc.Login(c.Request.Context(), &req); err != nil {
		if errors.Is(err, errInternalError) {
			responses.Fail(c, http.StatusInternalServerError, err.Error())
		} else if errors.Is(err, errInvalidCredentials) {
			responses.Fail(c, http.StatusUnauthorized, err.Error())
		} else {
			responses.Fail(c, http.StatusBadRequest, err.Error())
		}
		responses.Success(c, http.StatusOK, gin.H{"token": ss})
		return
	}
}
