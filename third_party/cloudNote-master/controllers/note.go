package controllers

import (
	"encoding/json"
	"services/note"
	"utils"
	//	"models"
	//	"strconv"
	//	"time"

	//	"comm"

	"github.com/astaxie/beego"
	//	"github.com/garyburd/redigo/redis"
)

type NoteController struct {
	BaseController
}

// Create new note
func (c *NoteController) Create() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := note.Create(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()
	//	var note models.Note
	//	if _, ok := Ob["Path"]; ok {
	//		note.CategoryId = Ob["Path"].(string)
	//		data := comm.GetDataAllFromRedis(UserId+note.CategoryId, UserId)
	//		maps, _ := redis.StringMap(data, err)
	//		if maps != nil {
	//			result := &models.Category{}
	//			err = FillStruct(maps, result)
	//			if err != nil {
	//				beego.Debug(err)
	//			} else {
	//				if result.IsDelete != 0 {
	//					Code = CATEGORY_ALREADY_DELETE_CODE
	//					Message = CATEGORY_ALREADY_DELETE
	//					c.ReturnJson(test)
	//					return
	//				}
	//			}
	//		} else {
	//			Code = CATEGORYID_WRONG_CODE
	//			Message = CATEGORYID_WRONG_MESSAGE
	//			c.ReturnJson(test)
	//			return
	//		}
	//	} else {
	//		c.InvalidParameter(test)
	//		return
	//	}
	//	if _, ok := Ob["Author"]; ok {
	//		note.Author = Ob["Author"].(string)
	//	}

	//	if _, ok := Ob["Title"]; ok {
	//		note.Title = Ob["Title"].(string)
	//	}

	//	if _, ok := Ob["Summary"]; ok {
	//		note.Summary = Ob["Summary"].(string)
	//	}

	//	if _, ok := Ob["Content"]; ok {
	//		note.Content = Ob["Content"].(string)
	//	} else {
	//		c.InvalidParameter(test)
	//		return
	//	}
	//	if _, ok := Ob["AttachmentList"]; ok {
	//		note.Attachment = Ob["AttachmentList"].(string)
	//	}
	//	if _, ok := Ob["CreateTime"]; ok {
	//		note.CreateTime = Ob["CreateTime"].(string)
	//	} else {
	//		CreateTime := int(time.Now().Unix())
	//		note.CreateTime = strconv.Itoa(CreateTime)
	//	}
	//	note.NoteId = "/4AF64012E9864C/123TDE3DC"
	//	data := StructToMap(note)

	//	beego.Debug(note)
	//	redisKey := UserId + note.NoteId
	//	Indexes := UserId + "note"
	//	err = comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	//	//	err := models.AddNote(&note)
	//	if CheckErr(err) {
	//		test.Path = "/4AF64012E9864C/123TDE3DC"
	//	}

	//	c.ReturnJson(test)
}

//Update note message
func (c *NoteController) Update() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := note.Create(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()

}

//Delete note
func (c *NoteController) Delete() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := note.Delete(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()

}

//Move note to next category
func (c *NoteController) Move() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := note.Move(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()

}
