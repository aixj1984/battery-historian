package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"

	"github.com/astaxie/beego/orm"
)

type Note struct {
	NoteId     string `orm:"column(noteId);pk"`
	CategoryId string `orm:"column(categoryId);size(50)"`
	Summary    string `orm:"column(summary);size(255);null"`
	CreateTime string `orm:"column(createTime);size(50);null"`
	Author     string `orm:"column(author);size(50);null"`
	Title      string `orm:"column(title);size(255);null"`
	Size       int    `orm:"column(size);null"`
	UpdateTime string `orm:"column(updateTime);size(50);null"`
	Content    string `orm:"column(content);null"`
	Attachment string `orm:"column(attachment);null"`
	IsDelete   int    `orm:"column(isDelete)"`
}

func (t *Note) TableName() string {
	return "note"
}

func init() {
	orm.RegisterModel(new(Note))
}
func GetNoteList(path string) (ch []Note, err error) {
	o := orm.NewOrm()
	var maps []Note
	sql := "SELECT * from note WHERE note.categoryId=? AND isDelete=0"
	num, err := o.Raw(sql, path).QueryRows(&maps)
	if num > 0 && err == nil {
		return maps, nil
	}
	return nil, err
}
func GetNoteById(path string, isDelete bool) (istrue bool, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	var sql string
	if isDelete {
		sql = "SELECT *  from note WHERE id=? AND isDelete=0"
	} else {
		sql = "SELECT *  from note WHERE id=? AND isDelete!=0"
	}
	num, err := o.Raw(sql, path).Values(&maps)
	if num > 0 && err == nil {
		return true, nil
	}
	return false, err
}
func GetOneNoteById(path string) (v Note, err error) {
	o := orm.NewOrm()
	var maps []Note
	var sql string
	sql = "SELECT *  from note WHERE id=? AND isDelete=0"

	num, err := o.Raw(sql, path).QueryRows(&maps)
	if num > 0 && err == nil {
		return maps[0], nil
	}
	return maps[0], err
}
func AddNote(v *Note) (err error) {
	o := orm.NewOrm()
	sql := "INSERT note VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	_, err = o.Raw(sql, v.NoteId, v.CategoryId, v.Summary, v.CreateTime, v.Author, v.Title, v.Size, v.UpdateTime, v.Content, v.Attachment, v.IsDelete).Exec()
	return err
}
func UpdateNote(v *Note) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE note  Set note.summary=?,note.author=?,note.title=?,note.size=?,note.updateTime=?,note.content=?,note.attachment=? WHERE noteId = ? AND isDelete==0"
	_, err = o.Raw(sql, v.Summary, v.Author, v.Title, v.Size, v.UpdateTime, v.Content, v.Attachment, v.NoteId).Exec()
	return err
}
func DeleteNote(path, updateTime string) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE note  Set note.isDelete=1,note.updateTime=? WHERE noteId = ? "
	_, err = o.Raw(sql, updateTime, path).Exec()

	return err
}
func DeleteNoteById(path string) (err error) {
	o := orm.NewOrm()
	sql := "DELETE FROM note WHERE noteId = ? "
	_, err = o.Raw(sql, path).Exec()

	return err
}
func RecoverNote(path string) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE note SET note.isDelete=0 WHERE noteId = ? "
	_, err = o.Raw(sql, path).Exec()
	return err
}
func MoveNote(path, categoryId string) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE note  Set note.categoryId=? WHERE noteId = ?"
	_, err = o.Raw(sql, categoryId, path).Exec()
	return err
}
func GetNoteDelete() ([]Note, error) {
	o := orm.NewOrm()
	var ch []Note
	num, err := o.QueryTable(new(Note)).Filter("isDelete", 1).All(&ch)
	if num > 0 && err == nil {
		return ch, nil
	}
	return nil, err
}
