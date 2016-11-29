package controllers

import (
	"encoding/json"
	"services/category"
	"utils"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	BaseController
}

//All get all category message
func (c *CategoryController) All() {
	cates := category.GetCategory()
	c.Data["json"] = cates
	c.ServeJSON()
}

//NoteList get one category's all note message
func (c *CategoryController) NoteList() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := category.GetNoteList(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()
	//	path := Ob["Path"].(string)
	//	data := comm.GetDataAllFromRedis(UserId+path, UserId)
	//	maps, _ := redis.StringMap(data, err)
	//	var notes []NoteList

	//	if maps == nil {
	//		Code = CATEGORYID_WRONG_CODE
	//		Message = CATEGORYID_WRONG_MESSAGE
	//		c.ReturnJson(notes)
	//		return
	//	}
	//	notes = GetAllNote(UserId+"note", path, UserId)
	//	var notes []NoteList
	//	note, err := models.GetNoteList(path)
	//	if len(note) == 0 {
	//		Code = CATEGORYID_WRONG_CODE
	//		Message = CATEGORYID_WRONG_MESSAGE
	//		c.ReturnJson(notes)
	//		return
	//	}
	//	if CheckErr(err) {
	//		for i := 0; i < len(note); i++ {
	//			var notelist NoteList
	//			notelist.Path = note[i].NoteId
	//			notelist.Summary = note[i].Summary
	//			notelist.Author = note[i].Author
	//			notelist.Title = note[i].Title
	//			attachlist := strings.Split(note[i].Attachment, ",")
	//			for i := 0; i < len(attachlist); i++ {
	//				var attach Attachment
	//				attachId, _ := strconv.Atoi(attachlist[i])
	//				attachment, err := models.GetAttachmentById(attachId)
	//				if err == nil {
	//					attach.Name = attachment.Name
	//					attach.Url = attachment.Url
	//					notelist.AttachmentList = append(notelist.AttachmentList, attach)
	//				}
	//			}

	//			notes = append(notes, notelist)
	//		}
	//	}
	//	beego.Debug(notes)
	//	c.ReturnJson(notes)
}

//Create create a new category
func (c *CategoryController) Create() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := category.Create(Ob)
	c.Data["json"] = result
	c.ServeJSON()
	//	var category models.Category
	//	if _, ok := Ob["CategoryName"]; ok {
	//		category.CategoryName = Ob["CategoryName"].(string)
	//	} else {
	//		c.InvalidParameter(test)
	//	}
	//	if _, ok := Ob["CreateTime"]; ok {
	//		createTime := Ob["CreateTime"].(string)
	//		category.CreateTime = createTime
	//	} else {
	//		createTime := int(time.Now().Unix())
	//		category.CreateTime = strconv.Itoa(createTime)
	//	}
	//	category.Id = "/4AF64012E9864C"
	//	category.UserId = UserId
	//	data := StructToMap(category)
	//	beego.Debug(category)
	//	redisKey := UserId + category.Id
	//	Indexes := UserId + "category"
	//	comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	//	err := models.AddCategory(&category)
	//	if CheckErr(err) {
	//		test.Path = category.Id
	//	}
	//	c.ReturnJson(test)
}

//Update update one category message
func (c *CategoryController) Update() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := category.Update(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()
	//	var category models.Category
	//	if _, ok := Ob["Path"]; ok {
	//		categoryId := Ob["Path"].(string)
	//		data := comm.GetDataAllFromRedis(UserId+categoryId, UserId)
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

	//		cate, err := models.GetOneCategoryById(categoryId)

	//		if err == nil && len(cate.Id) != 0 {
	//			category = cate
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
	//	if _, ok := Ob["CategoryName"]; ok {
	//		category.CategoryName = Ob["CategoryName"].(string)
	//	} else {
	//		c.InvalidParameter(test)
	//		return
	//	}
	//	if _, ok := Ob["UpdateTime"]; ok {
	//		category.UpdateTime = Ob["UpdateTime"].(string)
	//	} else {
	//		UpdateTime := int(time.Now().Unix())
	//		category.UpdateTime = strconv.Itoa(UpdateTime)

	//	}
	//	data := StructToMap(category)
	//	beego.Debug(category)
	//	redisKey := UserId + category.Id
	//	Indexes := UserId + "category"
	//	err = comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	//	//	err = models.UpdateCategoryList(&category)
	//	if CheckErr(err) {
	//		test.Result = "success"
	//	}
	//	c.ReturnJson(test)
}

//Delete delete one category
func (c *CategoryController) Delete() {
	beego.Debug(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &Ob)
	result := category.Delete(Ob, utils.UserId)
	c.Data["json"] = result
	c.ServeJSON()
	//	var path, updateTime string
	//	var test ReturnResult
	//	var category models.Category
	//	var err error
	//	if _, ok := Ob["UpdateTime"]; ok {
	//		updateTime = Ob["UpdateTime"].(string)
	//	}
	//	if _, ok := Ob["Path"]; ok {
	//		path = Ob["Path"].(string)
	//		data := comm.GetDataAllFromRedis(UserId+path, UserId)
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
	//		isNull, _ := models.GetCategoryById(path, true)
	//		if !isNull {
	//			Code = CATEGORYID_WRONG_CODE
	//			Message = CATEGORYID_WRONG_MESSAGE
	//			c.ReturnJson(test)
	//			return
	//		} else {
	//			isDelete, _ := models.GetCategoryById(path, false)
	//			if isDelete {
	//				Code = CATEGORY_ALREADY_DELETE_CODE
	//				Message = CATEGORY_ALREADY_DELETE
	//				c.ReturnJson(test)
	//				return
	//			}
	//		}
	//	}

	//	err := models.UpdateCategory(path, updateTime)
	//	category.IsDelete = 1
	//	category.UpdateTime = updateTime
	//	data := StructToMap(category)
	//	redisKey := UserId + category.Id
	//	Indexes := UserId + "category"
	//	err = comm.PutDataToRedis(redisKey, Indexes, UserId, data)
	//	if CheckErr(err) {
	//		test.Result = "success"
	//	}

	//	c.ReturnJson(test)
}
