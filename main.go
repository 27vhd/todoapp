package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/27vhd/todoapp/models"
	"github.com/google/uuid"
)

var validPath = regexp.MustCompile("^(/|/add/|/toggle/[0-9]+|/delete/[0-9]+)$")
var templates = template.Must(template.ParseFiles("./temps/list-todos.html", "./temps/add-todos.html"))
var dataPath = "./datas/todos.json"

func listTodosHandle(w http.ResponseWriter, r *http.Request) {
	_, err := checkPath(w, r)
	if err != nil {
		return
	}

	err = templates.ExecuteTemplate(w, "list-todos.html", loadTodos())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := checkPath(w, r)
	if err != nil {
		return
	}

	idInt, err := strconv.Atoi(id[len("/toggle/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	models.ToggleTodos(w, r, loadTodos(), uint32(idInt), dataPath)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := checkPath(w, r)
	if err != nil {
		return
	}
	idInt, err := strconv.Atoi(id[len("/delete/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	models.DeleteTodos(w, r, loadTodos(), uint32(idInt), dataPath)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := checkPath(w, r)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = templates.ExecuteTemplate(w, "add-todos.html", loadTodos())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		text := r.FormValue("text")
		models.AddTodos(w, r, loadTodos(), dataPath, &models.Todo{ID: uuid.New().ID(), Text: text, Done: false})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func loadTodos() *models.Todos {
	todos, err := models.GetTodos(dataPath)
	if err != nil {
		todos = &models.Todos{Todos: []*models.Todo{}}
	}

	return todos
}

func checkPath(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid url")
	}

	return m[0], nil
}

func main() {
	http.HandleFunc("/toggle/", toggleTodoHandler)
	http.HandleFunc("/delete/", deleteTodoHandler)
	http.HandleFunc("/add/", addTodoHandler)
	http.HandleFunc("/", listTodosHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
