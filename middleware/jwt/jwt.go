package jwt

import (
	"github.com/Alter/blog/pkg/retcode"
	"github.com/Alter/blog/pkg/util"
	"github.com/Alter/blog/routers/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = retcode.SUCCESS
		var data interface{}

		token := c.Query("token")
		if token == "" {
			api.ResponseData(c, retcode.INVALID_PARAMS, data)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = retcode.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = retcode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}

		if code != retcode.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  retcode.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}


}
