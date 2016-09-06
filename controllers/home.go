package controllers

import _ "fmt"

type HomeController struct {
	BaseController
}

type MachineConfig struct {
	Name   string
	Ip     string
	User   string
	Port   int
	Dbinfo string
}

type Users struct {
	Email    string
	Password string
}

//重写Prepare，会在每个method方法前调用
// func (this *HomeController) Prepare() {
// 	this.headerFile = "include/header.html"
// 	this.footerFile = "include/footer.html"
// 	this.layoutFile = "include/layout/classic.html"
// 	this.sidebarFile = "include/sidebar/classic_sidebar.html"
// }

func (this *HomeController) Index() {
	//this.Ctx.WriteString("aaa")
	flag := this.CheckLogin()
	if flag {
		this.getMachineConfig()
		this.MyRender("home/view_machine.html")
	} else {
		this.LoginRender("home/view_register.html")
	}

}

func (this *HomeController) Register() {
	// this.Ctx.WriteString("aaa")
	email := this.GetString("email")
	// this.Ctx.WriteString(email)
	password := this.GetString("password")
	defer this.session.Close()
	db := this.session.DB("monshark") //数据库名称
	collection := db.C("users")       //如果该集合已经存在的话，直接返回

	//插入数据
	err := collection.Insert(&Users{email, password})

	if err != nil {
		panic(err)
	} else {
		this.Ctx.WriteString("注册成功！")
	}

}

func (this *HomeController) getMachineConfig() {
	machineConfigData := map[string]MachineConfig{
		"bi":        {Name: "bi", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"oa":        {Name: "oa", Ip: "s119.29.00.59", User: "00", Port: 10220},
		"rsdk-set":  {Name: "rsdk-set", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"bi2-agent": {Name: "bi2-agent", Ip: "s119.00.25.59", User: "00", Port: 10220},
	}

	this.Data["machineConfigData"] = machineConfigData
}
