package accenrky
import (
	gct "github.com/elvtechnology/gocqltable"
	"github.com/elvtechnology/gocqltable/recipes"
	"github.com/gocql/gocql"
)

var (
	KeySpace gct.Keyspace
	TableStruct struct {recipes.CRUD}
)

type EntityManage interface {
	FindEntity(query string, keys ...interface{})
	FindALL()
}
func (e Entity) FindEntity(query string, keys ...interface{}) {

}
func (e Entity) FindALL()  {

}

func Find(crud struct {recipes.CRUD}) interface{} {

	rowset, err := crud.List()
	if err != nil {
		panic(err)
	}
	return rowset
}

//--------------------------------------------------------------------
func init() {
	c := gocql.NewCluster("127.0.0.1")

	s, err := c.CreateSession()
	if err != nil {
		panic(err)
	}

	gct.SetDefaultSession(s)

	KeySpace = gct.NewKeyspace("gocqltable_test")
}