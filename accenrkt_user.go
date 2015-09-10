package accenrky
import (
	"github.com/elvtechnology/gocqltable/recipes"
	"log"
)


func initCrud() struct{recipes.CRUD} {

	// Let's define and instantiate a table object for our user table
	userCrud := struct {
		recipes.CRUD // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			KeySpace.NewTable(
				"users",           // The table name
				[]string{"email"}, // Row keys
				nil,               // Range keys
				User{},            // We pass an instance of the user struct that will be used as a type template during fetches.
			),
		},
	}
	return userCrud
}

func (u User) FindALL() interface{} {

	rowset := Find(initCrud())
	users := rowset.([]*User)
	return users

}

func (u User) FindUserByEmail (email string) interface{} {

	// You can also fetch a single row, obviously
	row, err := initCrud().Get(email)
	if err != nil {
		log.Fatalln(err)
	}
	user := row.(*User)

	return user
}