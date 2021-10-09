package main

import (
	_ "appointy_api/routers"

	"github.com/astaxie/beego"
	"github.com/revel/revel"
	"github.com/kyawmyintthein/revel_mgo"
)

func main() {
	Config, err := revel.LoadConfig("app.conf")
    if err != nil || Config == nil {
        log.Fatalf("%+v",err)
    }
    mongodb.MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
    mongodb.PATH,_ = revel.Config.String("mongo.path")
    mongodb.DBNAME, _ = revel.Config.String("mongo.database")
    mongodb.CheckAndInitServiceConnection()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	
}
