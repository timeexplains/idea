package controller

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (aaa *MainController) Get(){
	//mydata :=`{"Website":"genius.com","Email":"email.com"}`
	//aaa.Data["json"]=mydata
	mydata :=`[{ "message": "Foo" },{ "message"": "Bar" }]`
	aaa.Data["items"] = mydata

	aaa.TplName = "index.html"
//	aaa.ServeJSON()
}

