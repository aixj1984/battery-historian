package attachment

import (
	"comm"
	"models"
	"utils"

	"github.com/astaxie/beego"
)

//Upload upload a new attachment
func Upload(fileType, fileName, UserId string) *utils.Response {
	var attach models.Attachment
	attach.Name = fileName
	attach.Type = fileType
	url := beego.AppConfig.String("download") + fileName
	attach.Url = url
	attach.Id = UserId + fileName
	data := utils.StructToMap(attach)
	redisKey := utils.UserId + attach.Id
	Indexes := utils.UserId + "attachment"
	err := comm.PutDataToRedis(redisKey, Indexes, utils.UserId, data)
	if err == nil {
		t := struct {
			Path string
		}{url}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}
