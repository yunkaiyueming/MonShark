package controllers

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/yunkaiyueming/MonShark/helpers"
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

//重写父类Prepare
func (this *UserController) Prepare() {
	fmt.Println("user controller prepare")

	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/main.html"
	this.sidebarFile = ""

	this.ConnMongoDB()
}

func (this *UserController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")
	action := this.GetString("action")
	if action == "" {
		this.MyRender("user/view_login.html")
		return
	}

	defer this.CloseMongoDB()
	collection := this.GetCol()

	result := User{}
	err := collection.Find(bson.M{"email": email}).One(&result)
	helpers.CheckError(err)

	if result.Password == password {
		this.SetSession("email", email)
		this.MyRedirect("home/index", 302)
	} else {
		this.MyRender("user/view_login.html")
	}
}

func (this *UserController) LogOut() {
	this.DelSession("email")
	this.MyRender("user/view_login.html")
}

func (this *UserController) Register() {
	this.MyRender("user/view_register.html")
}

func (this *UserController) DoRegister() {
	email := this.GetString("email")
	password := this.GetString("password")
	collection := this.GetCol()
	defer this.CloseMongoDB()

	err := collection.Insert(&User{bson.NewObjectId(), email, password})
	helpers.CheckError(err)

	this.SetSession("email", email)
	this.MyRedirect("home/index", 302)
}
