package accenrky
import "fmt"

var currentId int
var todos Todos

func init() {
	ReportCreateTodo(Todo{Name:"Write presentation"})
	ReportCreateTodo(Todo{Name:"Host meetup"})
}

func ReportFindTodo(id int) Todo  {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func ReportCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func ReportDestoryTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
