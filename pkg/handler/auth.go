package handler

import (
	"Resrik"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) sighUp(c *gin.Context) {
	var input Resrik.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type sighInInput struct {
	Username string `json:"username binding:required"`
	Password string `json:"password binding:required"`
}

func (h *Handler) sighIn(c *gin.Context) {
	var input sighInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
