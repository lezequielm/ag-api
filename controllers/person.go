package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lezequielm/ag-api/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type PersonController struct {
	beego.Controller
}

// @Title Create
// @Description create person
// @Param	body		body 	models.Person	true		"The Person content"
// @Success 200 {int} models.Person.id
// @Failure 403 body is empty
// @router / [post]
func (p *PersonController) Post() {
	var pr models.Person
	json.Unmarshal(p.Ctx.Input.RequestBody, &pr)
	o := orm.NewOrm()
	_, err := o.Insert(pr)
	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = pr
	}
	p.ServeJSON()
}

// @Title Get
// @Description find person by personId
// @Param	personId		path 	int	true		"the personid you want to get"
// @Success 200 {person} models.Person
// @Failure 403 :personId is empty
// @router /:personId [get]
func (p *PersonController) Get() {
	personId := p.Ctx.Input.Param(":personId")
	if personId != "" {
		intId, err := strconv.ParseInt(personId, 10, 32)
		pr, err := models.GetOnePerson(int(intId))
		if err != nil {
			p.Data["json"] = err.Error()
		} else {
			p.Data["json"] = pr
		}
	}
	p.ServeJSON()
}

// @Title GetAll
// @Description get all persons
// @Success 200 {person} models.Person
// @router / [get]
func (p *PersonController) GetAll() {
	prs, err := models.GetAllPersons()
	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = prs
	}
	p.ServeJSON()
}