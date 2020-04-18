package api

import (
	"net/http"
	"passport.xinfos.com/utils/errs"
	"passport.xinfos.com/pkg/logger"

	"github.com/gin-gonic/gin"
)

type response struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Next      string      `json:"next"`
}

//JSON - Public interface return
func JSON(c *gin.Context, httpCode int, data interface{}, err error) {
	if err != nil {
		logger.Error(err.Error())
	}
	c.JSON(http.StatusOK, response{
		Code: httpCode,
		Msg:  errs.ErrorMsg[httpCode],
		Data: data,
		Next: "",
	})

	return
}
