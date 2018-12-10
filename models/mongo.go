package models

import (
	//"fmt"
	//"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommonData struct {
	Source  string `bson:"source"` //
	Brand   string `bson:"brand"`
	Series  string `bson:"series"`
	Content string `bson:"content"`
}

const (
	MongoDBHosts = "47.96.184.66:3717"
	AuthDatabase = "crawler"
	AuthUserName = "ugc"
	AuthPassword = "a1b2c3d4"
)
const URL string = "47.96.184.66:3717"

var c *mgo.Collection
var session *mgo.Session

func init() {
	session, _ = mgo.Dial(URL)
	defer session.Close()
	//切换到数据库
	db := session.DB("crawler")
	//切换到collection
	c = db.C("public_praise")
}

/*
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err)   //终止运行
		}
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

//获取collection对象
func witchCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}
*/

func FindData(data map[string]interface{}) []CommonData {
	var commonData []CommonData

	c.Find(bson.M{"k_source": data["k_source"], "k_c_set": data["k_c_set"], "k_c_brand": data["k_c_brand"]}).One(&commonData)
	value := commonData
	return value
}

func CountData(data map[string]interface{}) (count int) {

	c.Find(bson.M{"k_source": data["k_source"], "k_c_set": data["k_c_set"], "k_c_brand": data["k_c_brand"]}).Count()
	return
}
