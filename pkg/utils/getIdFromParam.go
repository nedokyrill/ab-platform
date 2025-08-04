package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIdFromParam(param string, c *gin.Context) (uuid.UUID, bool) {
	str := c.Param(param)
	if str == "" {
		return uuid.Nil, false
	}

	parsedParam, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, false
	}

	return parsedParam, true
}
