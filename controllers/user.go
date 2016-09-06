package controllers

import (
	_ "fmt"

	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	BaseController
}

const DBNAME = "MonShark"
const CollectionName = "user"

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	JoinAt    time.Time     `bson:"joned_at"`
	Interests []string      `bson:"interests"`
}

func (this *UserController) GetDb() *mgo.Database {
	return this.mgoSession.DB(DBNAME)
}

func (this *UserController) GetCol() *mgo.Collection {
	return this.GetDb().C(CollectionName)
}

func (this *UserController) InsertTest() {
	data := &User{
		Id_:       bson.NewObjectId(),
		Name:      "Jimmy Kuu",
		Age:       33,
		JoinAt:    time.Now(),
		Interests: []string{"Develop", "Movie"},
	}

	err := this.GetCol().Insert(data)
	if err != nil {
		panic(err)
	}
}
