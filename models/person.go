package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Person struct {
	Id int `orm:"pk"`
	firstName string `json:"first_name"`
	lastName string `json:"last_name"`
	createAt time.Time `json:"create_at"`
	modifyAt time.Time `json:"modify_at"`
	enabled bool `json:"enabled"`
}

func init() {
	orm.RegisterModel(new(Person))
}

func GetOnePerson(id int) (Person, error) {
	o := orm.NewOrm()
	o.Using("default")
	p := Person{Id: id}
	err := o.Read(&p)
	return p, err
}

func GetAllPersons() ([]*Person, error) {
	o := orm.NewOrm()
	o.Using("default")
	var persons []*Person
	q := o.QueryTable(new(Person))
	num, err := q.All(&persons)
	fmt.Println(fmt.Printf(`results %d`, num))
	return persons, err
}