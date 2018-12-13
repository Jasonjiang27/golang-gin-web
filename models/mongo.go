package models

import (
	//"github.com/astaxie/beego/config"
	//"fmt"
	"time"
	"log"
	//"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type CommonData struct {
	Source  string `bson:"k_source"` //
	Brand   string `bson:"k_c_brand"`
	Series  string `bson:"k_c_set"`
	Content string `bson:"k_content"`
}

const URL string = "mongodb://ugc:a1b2c3d4@47.96.184.66:3717/crawler"

var c *mgo.Collection
var database *mgo.Database

func init() {
	
	dialInfo := &mgo.DialInfo{
		Addrs:		[]string{"47.96.184.66:3717"},
		Direct:		false,
		Database:	"crawler",
		Username:	"ugc",
		Password:	"a1b2c3d4",
		Timeout:	time.Second * 40,
		PoolLimit:	4096,

	}
	session, err := mgo.DialWithInfo(dialInfo)
	database = session.DB("crawler")
	
	//session, err := mgo.Dial(URL)

	if err != nil {
		log.Println(err.Error())
	}
	

	//session.SetMode(mgo.Monotonic, true)
	//使用指定的数据库
	//database = session.DB(config.Database)
}

func GetDataBase() *mgo.Database {
    return database
}

func GetErrNotFound() error {
    return mgo.ErrNotFound
}


/*
func GetDataSource()([]string, error) {
	var source []string
	con := GetDataBase().C("public_praise")
	if err := con.Find(bson.M{}).Distinct("k_source", &source); err != nil {
		if err.Error() != GetErrNotFound().Error() {
            return source, err
        }
	}
	return source, nil
}
*/
//mongo表中查询较慢，直接写死
func GetDataSource()([]string,error) {
	var err error
	source := []string{"pcauto", "autohome", "bitauto" ,"xcar", "sina", "12365auto", "qiche365"}
	return source, err
}

func GetBrands(maps map[string]interface{}) ([]string, error) {
	//补充mongo查询的数据
	var brand []string
	con := GetDataBase().C("car_brand")
	if err := con.Find(bson.M{"source":maps["k_source"]}).Distinct("name", &brand); err != nil {
		if err.Error() != GetErrNotFound().Error() {
            return brand, err
        }
	}
	return brand, nil
}

func GetSeries(maps map[string]interface{}) ([]string, error){
	var series []string
	con := GetDataBase().C("car_series")
	if err := con.Find(bson.M{"source":maps["k_source"],"brandName":maps["k_c_brand"]}).Distinct("name", &series); err != nil {
		if err.Error() != GetErrNotFound().Error() {
            return series, err
        }
	}
	return series, nil
}

/*
func init() {
	session, _ := mgo.DialWithTimeout(URL, 10 * time.Second)
	
	//切换到数据库
	db := session.DB("crawler")
	//切换到collection
	c = db.C("public_praise")
}
*/
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

func FindData(data map[string]interface{}) ([]CommonData, error) {
	var commonData []CommonData
	con := GetDataBase().C("public_praise")
	k_source := data["k_source"].(string)
	k_c_set := data["k_c_set"].(string)
	k_c_brand := data["k_c_brand"].(string)
	if err := con.Find(bson.M{"k_source": k_source, "k_c_set": k_c_set, "k_c_brand":  k_c_brand}).All(&commonData); err != nil {
		if err.Error() != GetErrNotFound().Error() {
            return commonData, err
        }
	}
	
	
	return commonData, nil
}


func CountData(data map[string]interface{}) (count int) {
	con := GetDataBase().C("public_praise")
	k_source := data["k_source"].(string)
	k_c_set := data["k_c_set"].(string)
	k_c_brand := data["k_c_brand"].(string)
	count, err := con.Find(bson.M{"k_source": k_source, "k_c_set": k_c_set, "k_c_brand": k_c_brand}).Count()
	if err != nil {
		panic(err)
	}
	return
}

func DbClose() {
	defer session.Close()
 }
