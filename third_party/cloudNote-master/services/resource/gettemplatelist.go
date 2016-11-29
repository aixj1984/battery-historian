package resource

import (
	"comm"
	"models"
	"utils"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//GetTemplateList get all template
func GetTemplateList() *utils.Response {
	cates := GetAttachment(utils.UserId+"attachment", "template", utils.UserId)
	return utils.NewResponse(0, "", cates)
}

//GetAllTemplate get all template
func GetAttachment(attachmentkey, attachType, userId string) []utils.Attachment {
	attachment := comm.GetDataKeyFromRedis(attachmentkey, userId)
	var attachs []utils.Attachment
	for i := 0; i < len(attachment); i++ {
		var attach utils.Attachment
		data := comm.GetDataAllFromRedis(attachment[i], userId)
		maps, _ := redis.StringMap(data, nil)
		if maps != nil {
			result := &models.Attachment{}
			err := utils.FillStruct(maps, result)
			if err != nil {
				beego.Debug(err)
			} else {
				if result.Type == attachType {
					attach.Name = result.Name
					attach.Url = result.Url
					attachs = append(attachs, attach)
				}

			}
		}
	}
	return attachs
}
