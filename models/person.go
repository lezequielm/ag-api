package models

import "github.com/astaxie/beego/orm"

type Person struct {
	id int64
	firstName string
	lastName string
	createAt orm.DateTimeField
	modifyAt orm.DateTimeField
	enabled bool
}