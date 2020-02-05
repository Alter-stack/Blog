package api

import (
	"github.com/Alter/blog/models"
	"github.com/Alter/blog/pkg/logging"
	"github.com/Alter/blog/pkg/retcode"
	"github.com/Alter/blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := retcode.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if !isExist {
			ResponseData(c, retcode.ERROR_AUTH, data)
			return
		}

		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = retcode.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = retcode.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	ResponseData(c, code, data)
}