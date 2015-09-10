package accenrky
import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"html"
	"io/ioutil"
	"io"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go.", html.EscapeString(r.URL.Path))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var u User
	users :=  u.FindALL()
	if err := json.NewEncoder(w).Encode(users); err != nil {panic(err)}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userEmials := vars["emails"]
	var u User
	user := u.FindUserByEmail(userEmials).(*User)
	fmt.Print(user)
	if user.Email != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(jsonErr{Code:http.StatusNotFound, Text:"Not Found the go"}); err != nil {
		panic(err)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request)  {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := ReportCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}