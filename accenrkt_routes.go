package accenrky
import (
	"net/http"
)

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserIndex",
		"GET",
		"/user",
		UserIndex,
	},
	Route{
		"UserShow",
		"GET",
		"/user/{emails}",
		UserShow,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreate,
	},
}
type Route struct {
	Name      	string
	Method    	string
	Pattren  	string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


