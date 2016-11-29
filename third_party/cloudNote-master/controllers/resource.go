package controllers

import (
	"services/resource"
)

type ResourceController struct {
	BaseController
}

//TemplateList get note template
func (c *ResourceController) TemplateList() {
	result := resource.GetTemplateList()
	c.Data["json"] = result
	c.ServeJSON()
	//	var tests []Attachment
	//	template, err := models.GetTemplate()
	//	if CheckErr(err) {
	//		for i := 0; i < len(template); i++ {
	//			var attach Attachment
	//			attach.Url = template[i].Url
	//			attach.Name = template[i].Name
	//			tests = append(tests, attach)
	//		}
	//	}
	//	c.ReturnJson(tests)
}

//PaperList get note paper
func (c *ResourceController) Paperlist() {
	result := resource.GetPaperList()
	c.Data["json"] = result
	c.ServeJSON()
	//	beego.Debug(string(c.Ctx.Input.RequestBody))
	//	var tests []Attachment
	//	template, err := models.GetPaper()
	//	if CheckErr(err) {
	//		for i := 0; i < len(template); i++ {
	//			var attach Attachment
	//			attach.Url = template[i].Url
	//			attach.Name = template[i].Name
	//			tests = append(tests, attach)
	//		}
	//	}
	//	c.ReturnJson(tests)
}
