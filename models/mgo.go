package models

import (
	"github.com/yunkaiyueming/MonShark/helpers"
	"gopkg.in/mgo.v2"
)

func GetDbConn(URL string) *mgo.Session {
	session, err := mgo.Dial(URL) //连接数据库
	helpers.CheckError(err)

	session.SetMode(mgo.Monotonic, true)
	return session
}
