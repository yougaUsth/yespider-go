package utils

import (
	"gopkg.in/mgo.v2"
	"log"
)


const gfsPrefix = "fs"


func NewMongoConnection(url, db, collection string) MongoConnection{
	mongo := MongoConnection{
		URL:url,
		DatabaseName:db,
		CollectionName:collection,
	}
	if err := mongo.Connection(); err != nil {
		panic(err.Error())
	}
	return mongo
}


type MongoConnection struct {
	URL            string
	DatabaseName   string
	CollectionName string
	Session        *mgo.Session
}

// 创建新的Session 与mongo 进行通信
func (m *MongoConnection) Connection() error{
	s, e := mgo.Dial(m.URL)
	if e != nil{
		return e
	}
	// 读从库
	//s.SetMode(mgo.Monotonic, true)
	m.Session = s
	return nil
}

func (m *MongoConnection) ensureConnected() (err error){
	defer func() {
		if r := recover(); r == nil{
			err := m.Connection()
			if err != nil{
				log.Fatalln(err.Error())
			}
		}
	}()
	err = m.Session.Ping()
	return
}


func (m *MongoConnection) getDB() *mgo.Database{
	return m.Session.DB(m.DatabaseName)
}

func (m *MongoConnection) getGFS() *mgo.GridFS {
	return m.getDB().GridFS(gfsPrefix)
}

func (m *MongoConnection) getColl() *mgo.Collection{
	return m.getDB().C(m.CollectionName)
}

func (m *MongoConnection) Find(query , selector interface{}) (*mgo.Query, error){
	return m.getColl().Find(query).Select(selector), nil
}
