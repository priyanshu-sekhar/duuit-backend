package main

import (
	_ "duuit/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		logs.GetLogger("ORM").Println("dev mode")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}