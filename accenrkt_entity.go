package accenrky
import (
	"github.com/gocql/gocql"
	"time"
)

type Entity struct {

}

type UserLog struct {
	Entity
	Email string     `json:"email"`// Row key part of our primary key
	Id    gocql.UUID `json:"id"`// Range key part of our primary key. Will use UUID version 1 (timeuuid).
	Data  int        `json:"date"`// The data we will log (in this case a integer value)
}

type User struct {
	Entity
	Email    string `json:"email"`// Our primary key
	Password string `json:"password"`     // Use Tags to rename fields
	Active   bool   `json:"active"` // If there are multiple tags, use `cql:""` to specify what the table column will be
	Created  time.Time `json:"created"`
}
