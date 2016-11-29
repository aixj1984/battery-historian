package category

import (
	"comm"
	"models"
	"utils"

	"strconv"
	"time"

	"github.com/astaxie/beego"
)

//Create create new category
func Create(Ob map[string]interface{}) *utils.Response {
	var category models.Category
	if _, ok := Ob["CategoryName"]; ok {
		category.CategoryName = Ob["CategoryName"].(string)
	} else {
		return utils.InvalidParameter()
	}
	if _, ok := Ob["CreateTime"]; ok {
		createTime := Ob["CreateTime"].(string)
		category.CreateTime = createTime
	} else {
		createTime := int(time.Now().Unix())
		category.CreateTime = strconv.Itoa(createTime)
	}
	category.Id = "/4AF64012E9864C"
	category.UserId = utils.UserId
	data := utils.StructToMap(category)
	beego.Debug(category)
	redisKey := utils.UserId + category.Id
	Indexes := utils.UserId + "category"
	err := comm.PutDataToRedis(redisKey, Indexes, utils.UserId, data)
	if err == nil {
		t := struct {
			Path string
		}{"/4AF64012E9864C"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}
