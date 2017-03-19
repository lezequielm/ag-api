package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Person struct {
	id int `orm:pk`
	firstName string `json:"first_name"`
	lastName string `json:"last_name"`
	createAt orm.DateTimeField `json:"create_at"`
	modifyAt orm.DateTimeField `json:"modify_at"`
	enabled bool `json:"enabled"`
}

func GetOnePerson(id int) (Person, error) {
	o := orm.NewOrm()
	o.Using("default")
	p := Person{id: id}
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