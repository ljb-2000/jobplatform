package dao

import (
	"github.com/astaxie/beego/orm"

)


func Db_insert(model interface{}) (int64, error) {
	o := orm.NewOrm()
	uid,err:=o.Insert(model)
	return uid,err
}