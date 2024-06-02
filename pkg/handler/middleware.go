package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid header")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "userId is not found")
		return 0, errors.New("userId not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "userId is invalid type")
		return 0, errors.New("userId not found")
	}

	return idInt, nil
}
