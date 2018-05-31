package Mongo

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type mtest struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Col1 string
	Col2 int
	Col3 bool
}

func Test_Mongo(t *testing.T) {
	t.Log("mongo test start")

	session := Mongo_getSession()
	defer session.Close()

	collection := session.DB("testDB").C("testCollection")
	collection.DropCollection()

	mtest1 := mtest{bson.NewObjectId(), "test1", 1, true}
	mtest2 := mtest{ID: bson.NewObjectId(), Col1: "test2", Col2: 2, Col3: false}
	//C
	if err := collection.Insert(&mtest1, &mtest2); err != nil {
		t.Log(err)
	}
	//R: read all
	result_all := []mtest{}
	if err := collection.Find(nil).All(&result_all); err != nil {
		t.Log(err)
	}
	t.Logf("size: %v", len(result_all))
	//R: read one
	result := mtest{}
	if err := collection.Find(bson.M{"col2": 2}).One(&result); err != nil {
		t.Log(err)
	}
	t.Log(result)

	//U
	//D
}
