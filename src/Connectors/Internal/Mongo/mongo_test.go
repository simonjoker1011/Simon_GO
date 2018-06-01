package Mongo

import (
	"Utils/Json"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type mtest struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"key"`
	Col1 string        `json:"att1"`
	Col2 int           `json:"att2"`
	Col3 bool          `json:"att3"`
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
	t.Logf("Insert records: \n %v", Utils.Marshal_To_String([]mtest{mtest1, mtest2}))
	if err := collection.Insert(&mtest1, &mtest2); err != nil {
		t.Log(err)
	}

	//R: read all
	result_all := []mtest{}
	if err := collection.Find(nil).All(&result_all); err != nil {
		t.Log(err)
	}
	t.Logf("read all size: %v", len(result_all))

	//R: read one
	col2value := 2
	t.Logf("read by col2: %v", col2value)
	result := mtest{}
	if err := collection.Find(bson.M{"col2": col2value}).One(&result); err != nil {
		t.Log(err)
	}
	t.Log(Utils.Marshal_To_String(result))

	//R: read by ID
	b, _ := mtest1.ID.MarshalJSON()
	t.Logf("read by id: %v", string(b[:]))
	if err := collection.Find(bson.M{"_id": mtest1.ID}).One(&result); err != nil {
		t.Log(err)
	}
	t.Log(Utils.Marshal_To_String(result))

	//U
	t.Log("update col3:")
	if err := collection.Update(bson.M{"col3": false}, bson.M{"col3": true}); err != nil {
		t.Log(err)
	}
	collection.Find(nil).All(&result_all)
	t.Log(Utils.Marshal_To_String(result_all))

	t.Log("update all col1")
	query := bson.M{"col3": true}
	update := bson.M{"$set": bson.M{"col1": "updated"}}
	info, err := collection.UpdateAll(query, update)
	if err != nil {
		t.Log(err)
	}
	t.Log(Utils.Marshal_To_String(*info))
	collection.Find(nil).All(&result_all)
	t.Log(Utils.Marshal_To_String(result_all))

	//D
	collection.Find(bson.M{"col2": 1}).All(&result_all)
	t.Log(Utils.Marshal_To_String(result_all))

	if err := collection.Remove(bson.M{"col2": 1}); err != nil {
		t.Log(err)
	}
	collection.Find(nil).All(&result_all)
	t.Log(Utils.Marshal_To_String(result_all))

}
