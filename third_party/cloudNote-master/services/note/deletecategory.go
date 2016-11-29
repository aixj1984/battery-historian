package note

import (
	"comm"
	"models"
	"utils"
)

//Delete delete one note
func Delete(Ob map[string]interface{}, UserId string) *utils.Response {
	var note models.Note
	var noteId, updateTime string
	if _, ok := Ob["UpdateTime"]; ok {
		updateTime = Ob["UpdateTime"].(string)
	}
	if _, ok := Ob["Path"]; ok {
		noteId = Ob["Path"].(string)
		isHave, result := CheckNote(UserId+noteId, UserId)
		if !isHave {
			return result
		}
	} else {
		utils.InvalidParameter()
	}
	note.IsDelete = 1
	note.UpdateTime = updateTime
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
