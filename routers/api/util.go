package api

import (
	"github.com/Alter/blog/pkg/retcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseData(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : retcode.GetMsg(code),
		"data" : data,
	})
	return
}
