package category

import (
	"comm"
	"models"
	"utils"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//GetCategory get all category message
func GetCategory() *utils.Response {
	cates := GetAllCategory(utils.UserId+"category", utils.UserId+"note", utils.UserId)
	return utils.NewResponse(0, "", cates)
}

//GetAllCategory get all category
func GetAllCategory(categoryKey, noteKey, userId string) []utils.Category {
	noteIds := comm.GetDataKeyFromRedis(categoryKey, userId)
	var cates []utils.Category
	for i := 0; i < len(noteIds); i++ {
		var cate utils.Category
		data := comm.GetDataAllFromRedis(noteIds[i], userId)
		maps, _ := redis.StringMap(data, nil)
		if maps != nil {
			result := &models.Category{}
			err := utils.FillStruct(maps, result)
			if err != nil {
				beego.Debug(err)
			} else {
				if result.IsDelete == 0 {
					cate.Name = result.CategoryName
					cate.CreateTime = result.CreateTime
					cate.UpdateTime = result.UpdateTime
					notes := GetAllNote(noteKey, result.Id, userId)
					cate.NoteCount = len(notes)
					cates = append(cates, cate)
				}

			}
		}
	}
	return cates
}
