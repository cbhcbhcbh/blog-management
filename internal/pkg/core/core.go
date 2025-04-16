package core

import (
	"blog/internal/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Code string `json:"code"`

	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
