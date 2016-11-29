package note

import (
	"comm"
	"models"
	"utils"

	"strconv"
	"time"
)

//Update update one note
func Update(Ob map[string]interface{}, UserId string) *utils.Response {
	var note models.Note
	if _, ok := Ob["Path"]; ok {
		noteId := Ob["Path"].(string)
		isHave, result := CheckNote(UserId+noteId, UserId)
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
	}
	if _, ok := Ob["AttachmentList"]; ok {
		note.Attachment = Ob["AttachmentList"].(string)
	}
	if _, ok := Ob["UpdateTime"]; ok {
		note.UpdateTime = Ob["UpdateTime"].(string)
	} else {
		UpdateTime := int(time.Now().Unix())
		note.UpdateTime = strconv.Itoa(UpdateTime)
	}
	note.NoteId = "/4AF64012E9864C/123TDE3DC"
	data := utils.StructToMap(note)
	redisKey := UserId + note.NoteId
	Indexes := UserId + "note"
	err := comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	if err == nil {
		t := struct {
			Path string
		}{"//4AF64012E9864C/123TDE3DC"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}
