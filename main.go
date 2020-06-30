package main

import (
	_ "github.com/yanluwei/qraphql-go/routers"
	"github.com/yanluwei/qraphql-go/schema"
	"net/http"
)

//func main() {
//	if beego.BConfig.RunMode == "dev" {
//		beego.BConfig.WebConfig.DirectoryIndex = true
//		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
//	}
//	beego.Run()
//
//
//}

func main() {

	http.Handle("/hello", schema.HelloHandler)
	http.Handle("/hello1", schema.Hello1Handler)
	http.ListenAndServe(":8080", nil)
}
