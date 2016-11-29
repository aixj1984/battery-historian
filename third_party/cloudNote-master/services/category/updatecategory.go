package category

import (
	"comm"
	"models"
	"utils"

	"strconv"
	"time"

	"github.com/astaxie/beego"
)

//Update Update one category
func Update(Ob map[string]interface{}, UserId string) *utils.Response {
	var category models.Category
	if _, ok := Ob["Path"]; ok {
		categoryId := Ob["Path"].(string)
		isHave, result := CheckCategoryId(UserId+categoryId, UserId)
		if !isHave {
			return result
		}
	} else {
		return utils.InvalidParameter()

	}
	if _, ok := Ob["CategoryName"]; ok {
		category.CategoryName = Ob["CategoryName"].(string)
	} else {
		return utils.InvalidParameter()
	}
	if _, ok := Ob["UpdateTime"]; ok {
		category.UpdateTime = Ob["UpdateTime"].(string)
	} else {
		UpdateTime := int(time.Now().Unix())
		category.UpdateTime = strconv.Itoa(UpdateTime)
	}
	data := utils.StructToMap(category)
	beego.Debug(category)
	redisKey := UserId + category.Id
	Indexes := UserId + "category"
	err := comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	if err == nil {
		t := struct {
			Result string
		}{"success"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}
