package main

import (
	_ "github.com/lezequielm/ag-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/lezequielm/ag-api/models"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err:=orm.RegisterDataBase("default", "sqlite3", "file:data.db")
	orm.RegisterModel(new(models.Person))
	name := "default"
	force := true
	verbose := true
	err = orm.RunSyncdb(name, force, verbose)
	if err!=nil{
		panic(err)
	}

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
