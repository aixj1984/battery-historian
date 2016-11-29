package note

import (
	"comm"
	"models"
	"services/category"
	"utils"

	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//Create create new note
func Create(Ob map[string]interface{}, UserId string) *utils.Response {
	var note models.Note
	if _, ok := Ob["Path"]; ok {
		note.CategoryId = Ob["Path"].(string)
		isHave, result := category.CheckCategoryId(UserId+note.CategoryId, UserId)
		if !isHave {
			return result
		}
	} else {
		return utils.InvalidParameter()
	}
	if _, ok := Ob["Author"]; ok {
		note.Author = Ob["Author"].(string)
	}

	if _, ok := Ob["Title"]; ok {
		note.Title = Ob["Title"].(string)
	}

	if _, ok := Ob["Summary"]; ok {
		note.Summary = Ob["Summary"].(string)
	}
	if _, ok := Ob["Content"]; ok {
		note.Content = Ob["Content"].(string)
	} else {
		return utils.InvalidParameter()
	}
	if _, ok := Ob["AttachmentList"]; ok {
		note.Attachment = Ob["AttachmentList"].(string)
	}
	if _, ok := Ob["CreateTime"]; ok {
		note.CreateTime = Ob["CreateTime"].(string)
	} else {
		CreateTime := int(time.Now().Unix())
		note.CreateTime = strconv.Itoa(CreateTime)
	}
	note.NoteId = "/4AF64012E9864C/123TDE3DC"
	data := utils.StructToMap(note)
	redisKey := UserId + note.NoteId
	Indexes := UserId + "note"
	err := comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	if err == nil {
		t := struct {
			Path string
		}{"/4AF64012E9864C/123TDE3DC"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}

func CheckNote(key, userId string) (bool, *utils.Response) {
	data := comm.GetDataAllFromRedis(key, userId)
	maps, _ := redis.StringMap(data, nil)
	if maps == nil {
		return false, utils.NewResponse(utils.NOTE_NOT_EXIST_CODE, utils.NOTE_NOT_EXIST, nil)
	} else {
		result := &models.Category{}
		err := utils.FillStruct(maps, result)
		if err != nil {
			beego.Debug(err)
		} else {
			if result.IsDelete != 0 {
				return false, utils.NewResponse(utils.NOTE_ALREADY_DELETE_CODE, utils.NOTE_ALREADY_DELETE, nil)
			}
		}
	}
	return true, nil
}
