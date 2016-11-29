package resource

import (
	"utils"
)

//GetPaperList get all template
func GetPaperList() *utils.Response {
	cates := GetAttachment(utils.UserId+"attachment", "paper", utils.UserId)
	return utils.NewResponse(0, "", cates)
}
