package note

import (
	"comm"
	"models"
	"services/category"
	"utils"
)

//Move move one note to other category
func Move(Ob map[string]interface{}, UserId string) *utils.Response {
	var noteId, categoryId string
	var note models.Note
	if _, ok := Ob["Path"]; ok {
		noteId = Ob["Path"].(string)
		isHave, result := CheckNote(UserId+noteId, UserId)
		if !isHave {
			return result
		}
	} else {
		utils.InvalidParameter()
	}
	if _, ok := Ob["GoalPath"]; ok {
		categoryId = Ob["GoalPath"].(string)
		isHave, result := category.CheckCategoryId(UserId+note.CategoryId, UserId)
		if !isHave {
			return result
		}
	} else {
		utils.InvalidParameter()
	}
	note.CategoryId = categoryId
	data := utils.StructToMap(note)
	redisKey := UserId + noteId
	Indexes := UserId + "note"
	err := comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	if err == nil {
		t := struct {
			Result string
		}{"success"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}
