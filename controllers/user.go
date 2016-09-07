package controllers

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	BaseController
}

type User struct {
	Id_      bson.ObjectId `bson:"_id"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

const DBNAME = "MonShark"
const CollectionName = "user"

func (this *UserController) GetDb() *mgo.Database {
	return this.mgoSession.DB(DBNAME)
}

func (this *UserController) GetCol() *mgo.Collection {
	return this.GetDb().C(CollectionName)
}

func (this *UserController) DoLogin() {
	email := this.GetString("email")
	password := this.GetString("password")

	defer this.CloseMongoDB()

	// db := this.mgoSession.DB(DBNAME)   //数据库名称
	// collection := db.C(CollectionName) //如果该集合已经存在的话，直接返回
	collection := this.GetCol()

	//****查询单条数据****
	result := User{}
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
	collection := this.GetCol()
	defer this.CloseMongoDB()
	//插入数据
	err := collection.Insert(&User{bson.NewObjectId(), email, password})

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
