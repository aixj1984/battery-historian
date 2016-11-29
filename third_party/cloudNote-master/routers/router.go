package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/V1.0",

		beego.NSNamespace("/Attachment",
			beego.NSRouter("/Upload", &controllers.AttachmentController{}, "post:Upload"),
			beego.NSRouter("/DownLoad", &controllers.AttachmentController{}, "get:DownLoad"),
		),
		beego.NSNamespace("/Category",
			beego.NSRouter("/All", &controllers.CategoryController{}, "get:All"),
			beego.NSRouter("/NoteList", &controllers.CategoryController{}, "post:NoteList"),
			beego.NSRouter("/Create", &controllers.CategoryController{}, "post:Create"),
			beego.NSRouter("/Update", &controllers.CategoryController{}, "post:Update"),
			beego.NSRouter("/Delete", &controllers.CategoryController{}, "post:Delete"),
		),

		beego.NSNamespace("/Note",
			beego.NSRouter("/Move", &controllers.NoteController{}, "post:Move"),
			beego.NSRouter("/Create", &controllers.NoteController{}, "post:Create"),
			beego.NSRouter("/Update", &controllers.NoteController{}, "post:Update"),
			beego.NSRouter("/Delete", &controllers.NoteController{}, "post:Delete"),
		),
		beego.NSNamespace("/Resource",
			beego.NSRouter("/TemplateList", &controllers.ResourceController{}, "get:TemplateList"),
			beego.NSRouter("/Paperlist", &controllers.ResourceController{}, "get:Paperlist"),
		),
		beego.NSNamespace("/Share",
			beego.NSRouter("/Publish", &controllers.ShareController{}, "post:Share"),
			beego.NSRouter("/Show/:id", &controllers.ShareController{}, "get:Show"),
		),
	)
	beego.AddNamespace(ns)
}
