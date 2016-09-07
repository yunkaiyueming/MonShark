package controllers

import "gopkg.in/mgo.v2/bson"

type UserController struct {
	BaseController
}

func (this *UserController) DoLogin() {
	email := this.GetString("email")
	password := this.GetString("password")

	defer this.CloseMongoDB()

	db := this.mgoSession.DB("monshark") //数据库名称
	collection := db.C("users")          //如果该集合已经存在的话，直接返回

	//****查询单条数据****
	result := Users{}
	err := collection.Find(bson.M{"email": email}).One(&result)

	if err != nil {
		panic(err)
	} else {
		if result.Password == password {
			this.SetSession("email", email)
			//this.Ctx.WriteString("登录成功！")
			this.getMachineConfig()
			this.MyRender("home/view_machine.html")
		}
	}

}

func (this *UserController) LogOut() {
	this.DelSession("email")
	this.LoginRender("home/view_welcome.html")
}

func (this *UserController) DoRegister() {
	email := this.GetString("email")
	password := this.GetString("password")
	db := this.mgoSession.DB("monshark") //数据库名称
	collection := db.C("users")          //如果该集合已经存在的话，直接返回
	defer this.CloseMongoDB()
	//插入数据
	err := collection.Insert(&Users{email, password})

	if err != nil {
		panic(err)
	} else {
		this.SetSession("email", email)
		this.getMachineConfig()
		this.MyRender("home/view_machine.html")
	}

}

func (this *UserController) Register() {
	this.LoginRender("home/view_register.html")
}
