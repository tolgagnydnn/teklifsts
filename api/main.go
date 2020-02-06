package main

import (
	_ "api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	corsHandler := cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://127.0.0.1:8080/"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
