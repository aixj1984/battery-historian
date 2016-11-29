package category

import (
	"comm"
	"models"
	"utils"

	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//Delete delete one category
func Delete(Ob map[string]interface{}, UserId string) *utils.Response {
	var path, updateTime string
	var category models.Category
	var err error
	if _, ok := Ob["UpdateTime"]; ok {
		updateTime = Ob["UpdateTime"].(string)
	} else {
		tm := int(time.Now().Unix())
		updateTime = strconv.Itoa(tm)
	}
	if _, ok := Ob["Path"]; ok {
		path = Ob["Path"].(string)
		isHave, result := CheckCategoryId(UserId+path, UserId)
		if !isHave {
			return result
		}
	}
	category.IsDelete = 1
	category.UpdateTime = updateTime
	data := utils.StructToMap(category)
	redisKey := UserId + category.Id
	Indexes := UserId + "category"
	err = comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	if err == nil {
		t := struct {
			Result string
		}{"success"}
		return utils.NewResponse(0, "", t)
	}
	return utils.NewResponse(utils.SYSTEM_CODE, err.Error(), nil)
}

func CheckCategoryId(key, userid string) (bool, *utils.Response) {
	data := comm.GetDataAllFromRedis(key, userid)
	maps, _ := redis.StringMap(data, nil)
	if maps == nil {
		return false, utils.NewResponse(utils.CATEGORYID_WRONG_CODE, utils.CATEGORYID_WRONG_MESSAGE, nil)
	} else {
		result := &models.Category{}
		err := utils.FillStruct(maps, result)
		if err != nil {
			beego.Debug(err)
		} else {
			if result.IsDelete != 0 {
				return false, utils.NewResponse(utils.CATEGORY_ALREADY_DELETE_CODE, utils.CATEGORY_ALREADY_DELETE, nil)
			}
		}
	}
	return true, nil

}
