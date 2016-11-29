package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id           string `orm:"column(id);pk"`
	CategoryName string `orm:"column(categoryName);size(255);null"`
	CreateTime   string `orm:"column(createTime);size(50);null"`
	UpdateTime   string `orm:"column(updateTime);size(50);null"`
	IsDelete     int    `orm:"column(isDelete);"`
	UserId       string `orm:"column(userId);size(255);null"`
}

func (t *Category) TableName() string {
	return "category"
}

func init() {
	orm.RegisterModel(new(Category))
}
func GetCategoryList() (ch []orm.Params, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := "SELECT id,categoryName,category.createTime,category.updateTime,Count(note.categoryId) as noteNum  from category LEFT JOIN note ON note.categoryId=id AND note.isDelete=0 WHERE category.isDelete=0 GROUP BY category.id"
	num, err := o.Raw(sql).Values(&maps)
	if num > 0 && err == nil {
		return maps, nil
	}
	return nil, err
}
func GetCategoryByNameAndUserId(name, userId string) (ch []orm.Params, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := "SELECT * FROM category WHERE categoryName=? AND userId=?"
	num, err := o.Raw(sql, name, userId).Values(&maps)
	if num > 0 && err == nil {
		return maps, nil
	}
	return nil, err
}
func GetOneCategoryById(path string) (v Category, err error) {
	o := orm.NewOrm()
	var maps []Category
	var sql string
	sql = "SELECT * from category WHERE id=? AND isDelete=0"

	num, err := o.Raw(sql, path).QueryRows(&maps)
	if num > 0 && err == nil {
		return maps[0], nil
	}
	return maps[0], err
}
func GetCategoryById(path string, isDelete bool) (istrue bool, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	var sql string
	if isDelete {
		sql = "SELECT *  from category WHERE id=? AND isDelete=0"
	} else {
		sql = "SELECT *  from category WHERE id=? AND isDelete!=0"
	}
	num, err := o.Raw(sql, path).Values(&maps)
	if num > 0 && err == nil {
		return true, nil
	}
	return false, err
}
func AddCategory(v *Category) (err error) {
	o := orm.NewOrm()
	sql := "INSERT category VALUES(?,?,?,?,?,?)"
	_, err = o.Raw(sql, v.Id, v.CategoryName, v.CreateTime, v.UpdateTime, v.IsDelete, v.UserId).Exec()
	return err
}
func UpdateCategoryList(v *Category) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE category  Set category.categoryName=?,category.updateTime=? WHERE Id = ?"
	_, err = o.Raw(sql, v.CategoryName, v.UpdateTime, v.Id).Exec()

	return err
}
func DeleteCategoryById(path string) (err error) {
	o := orm.NewOrm()
	sql := "DELETE FROM category WHERE id =?"
	_, err = o.Raw(sql, path).Exec()

	return err
}
func UpdateCategory(path, updateTime string) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE category  Set category.isDelete=1,category.updateTime=? WHERE Id = ?"
	_, err = o.Raw(sql, updateTime, path).Exec()
	if err == nil {
		sql := "UPDATE note  Set note.isDelete=1,note.updateTime=? WHERE note.categoryId = ?"
		_, err = o.Raw(sql, updateTime, path).Exec()
	}
	return err
}
func RecoverCategory(path string) error {
	o := orm.NewOrm()
	sql := "UPDATE category  Set category.isDelete=0 category.updateTime=? WHERE Id = ?"
	_, err := o.Raw(sql, path).Exec()

	return err
}
func GetCategoryDelete() ([]Category, error) {
	o := orm.NewOrm()
	var ch []Category
	num, err := o.QueryTable(new(Category)).Filter("isDelete", 1).All(&ch)
	if num > 0 && err == nil {
		return ch, nil
	}
	return nil, err
}
