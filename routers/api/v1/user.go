package v1

import (
	"github.com/Alter/blog/pkg/logging"
	"github.com/Alter/blog/pkg/retcode"
	"github.com/Alter/blog/pkg/setting"
	"github.com/Alter/blog/pkg/util"
	"github.com/Alter/blog/pkg/util/pbcodec"
	"github.com/Alter/blog/routers/api"
	service "github.com/Alter/blog/service/pb_service/protofile"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data := make(map[string]interface{})
	pbcodec.InitTransfer()
	readAllRequest := &service.ReadAllRequest{
		Token: "",
		Pageno: int64(util.GetPage(c)),
		Count: int64(setting.AppSetting.PageSize),
	}

	if err := pbcodec.G_transfer.SendMsg(int(service.MyMessage_ReadAllRequest), readAllRequest); err != nil {
		api.ResponseData(c, retcode.ERROR, data)
		logging.Error(err)
		return
	}
	var readAllResponse service.ReadAllResponse
	if err := pbcodec.G_transfer.ReadMsg(&readAllResponse); err != nil {
		api.ResponseData(c, retcode.ERROR, data)
		logging.Error(err)
		return
	}
	code := retcode.SUCCESS
	data["lists"] = readAllResponse.Users
	data["total"] = len(readAllResponse.Users)

	api.ResponseData(c, code, data)
}

func GetUser(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func EditUser(c *gin.Context) {

}
