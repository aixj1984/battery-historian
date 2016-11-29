package controllers

import (
	"services/attachment"
	"utils"

	"github.com/astaxie/beego"
)

type AttachmentController struct {
	BaseController
}

//Upload new attachment
func (c *AttachmentController) Upload() {
	var result *utils.Response
	fileType := c.Input().Get("fileType")
	f, h, err := c.GetFile("File")
	defer f.Close()
	if err != nil {
		beego.Debug("getfile err ", err)
	} else {
		err = c.SaveToFile("File", "./swagger/upload/"+h.Filename)
		if err == nil {
			result = attachment.Upload(fileType, h.Filename, utils.UserId)
		} else {
			beego.Debug(err)
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *AttachmentController) DownLoad() {
	fileName := c.Input().Get("id")
	c.Ctx.Output.Download("./swagger/upload" + fileName)
}
