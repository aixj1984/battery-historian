package controllers

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/astaxie/beego"
)

type ShareController struct {
	BaseController
}

type NoteMessage struct {
	Key               string
	Text              string
	Type              string
	Depth             int
	InlineStyleRanges []InlineStyleRange
	EntityRanges      []EntityRange
	Data              Datas
}
type Datas struct {
	Attach
	WrapKey    string
	BTextAlign string
}
type InlineStyleRange struct {
	Offset int
	Length int
	Style  string
}
type EntityRange struct {
}
type Attach struct {
	Src   string
	Type  string
	style interface{}
}

func (c *ShareController) Share() {
	//	beego.Debug(string(c.Ctx.Input.RequestBody))

	//	var test ReturnPath
	//	test.Path = "www.baidu.com"
	//	c.ReturnJson(test)
}
func (c *ShareController) Show() {
	//	path := c.Input().Get(":id")
	var noteMessage []NoteMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &noteMessage)
	beego.Debug(noteMessage)
	context := JSONToHtml(noteMessage)
	c.Ctx.WriteString(context)
}
func JSONToHtml([]NoteMessage) string {
	// 声明模板内容
	const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
    </body>
</html>`

	check := func(err error) {
		if err != nil {
			beego.Debug(err)
		}
	}

	// 创建一个新的模板，并且载入内容
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	// 定义传入到模板的数据，并在终端打印
	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	wr := bytes.NewBufferString("")
	err = t.Execute(wr, data)
	check(err)
	return wr.String()
}
