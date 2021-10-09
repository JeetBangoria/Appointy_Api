package main

import (
	_ "appointy_api/routers"
	"appointy_api/controllers"
	"net/http"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	
)

//var client *mongo.Client

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		httpRoute := httprouter.New()
		user := controllers.UserController(getSession())
		post := controllers.PostController(getSession())
		http.ListenAndServe("localhost:8080", httpRoute)
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.StaticDir["/users"] = user.CreateUser
		beego.BConfig.WebConfig.StaticDir["/users/:id"] = user.GetUser
		beego.BConfig.WebConfig.StaticDir["/posts"] = post.CreatePost
		beego.BConfig.WebConfig.StaticDir["/posts/:id"] = post.GetPost
		beego.BConfig.WebConfig.StaticDir["/posts/users/:id"] = post.GetAllUserPost
	}
	beego.Run()
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}


