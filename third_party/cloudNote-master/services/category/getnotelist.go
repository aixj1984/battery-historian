package category

import (
	"comm"
	"models"
	"strings"
	"utils"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

func GetNoteList(Ob map[string]interface{}, userId string) *utils.Response {
	var path string
	if _, ok := Ob["Path"]; ok {
		path = Ob["Path"].(string)
	} else {
		return utils.InvalidParameter()
	}
	isHave, result := CheckCategoryId(userId+path, userId)
	if !isHave {
		return result
	}
	return utils.NewResponse(0, "", GetAllNote(userId+"note", path, userId))
}

//GetAllNote get all note int one category
func GetAllNote(key, categoryId, userId string) []utils.NoteList {
	noteIds := comm.GetDataKeyFromRedis(key, userId)
	var notes []utils.NoteList
	for i := 0; i < len(noteIds); i++ {
		var note utils.NoteList
		data := comm.GetDataAllFromRedis(noteIds[i], userId)
		maps, _ := redis.StringMap(data, nil)
		if maps != nil {
			result := &models.Note{}
			err := utils.FillStruct(maps, result)
			if err != nil {
				beego.Debug(err)
			} else {
				if result.CategoryId == categoryId {
					note.Path = result.NoteId
					note.Summary = result.Summary
					note.Author = result.Author
					note.Title = result.Title
					attachlist := strings.Split(result.Attachment, ",")
					for i := 0; i < len(attachlist); i++ {
						var single utils.Attachment
						data := comm.GetDataAllFromRedis(attachlist[i], userId)
						maps, _ := redis.StringMap(data, nil)
						attachment := &models.Attachment{}
						err = utils.FillStruct(maps, result)
						if err == nil {
							single.Name = attachment.Name
							single.Url = attachment.Url
						}
						note.AttachmentList = append(note.AttachmentList, single)
					}
					notes = append(notes, note)
				}

			}
		}
	}
	return notes
}
