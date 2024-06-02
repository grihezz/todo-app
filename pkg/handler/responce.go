package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponce{message})
}
