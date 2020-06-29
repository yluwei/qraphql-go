package controllers

import "github.com/astaxie/beego"

type StudentController struct {
	beego.Controller
}

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s *StudentController) ListStudent() {
	student := Student{
		Id:   1,
		Name: "小明",
	}
	s.Data["json"] = student
	s.ServeJSON()
}
