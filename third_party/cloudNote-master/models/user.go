package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	UserId         string    `orm:"column(userId);pk"`
	UserName       string    `orm:"column(userName);size(255);null"`
	Password       string    `orm:"column(password);size(50)"`
	TotalSize      int       `orm:"column(totalSize);null"`
	UsedSize       int       `orm:"column(usedSize);type(datetime);null"`
	RegisterTime   time.Time `orm:"column(registerTime);type(datetime);null"`
	LastLoginTime  time.Time `orm:"column(lastLoginTime);type(datetime);null"`
	LastUpdateTime time.Time `orm:"column(lastUpdateTime);type(datetime);null"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
func GetOneUserById(path string) (v User, err error) {
	o := orm.NewOrm()
	var maps []User
	var sql string
	sql = "SELECT * from user WHERE userId=?"

	num, err := o.Raw(sql, path).QueryRows(&maps)
	if num > 0 && err == nil {
		return maps[0], nil
	}
	return maps[0], err
}
func AddUser(v *User) (err error) {
	o := orm.NewOrm()
	sql := "INSERT user VALUES(?,?,?,?,?,?,?,?)"
	_, err = o.Raw(sql, v.UserId, v.UserName, v.Password, v.TotalSize, v.UsedSize, v.RegisterTime, v.LastLoginTime, v.LastUpdateTime).Exec()
	return err
}

func UpdateUser(v *User) (err error) {
	o := orm.NewOrm()
	sql := "UPDATE user  Set userName=?,totalSize=?,usedSize=?,registerTime=?,lastLoginTime=?,lastUpdateTime=? WHERE userId = ?"
	_, err = o.Raw(sql, v.UserName, v.TotalSize, v.UsedSize, v.RegisterTime, v.LastLoginTime, v.LastUpdateTime, v.UserId).Exec()

	return err
}
