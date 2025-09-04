package models

import (
	"encoding/json"
	"net/http"
	"os"
)

type Todos struct {
	Todos []*Todo `json:"todos"`
}

func GetTodos(dataPath string) (*Todos, error) {
	var todos *Todos

	data, err := os.ReadFile(dataPath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func ToggleTodos(w http.ResponseWriter, r *http.Request, todos *Todos, id uint32, dataPath string) {
	for _, todo := range todos.Todos {
		if todo.ID == id {
			todo.Done = !todo.Done
		}
	}

	data, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	os.WriteFile(dataPath, data, 0600)

	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteTodos(w http.ResponseWriter, r *http.Request, todos *Todos, id uint32, dataPath string) {
	for i, todo := range todos.Todos {
		if todo.ID == id {
			todos.Todos = append(todos.Todos[:i], todos.Todos[i+1:]...)
		}
	}

	data, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	os.WriteFile(dataPath, data, 0600)

	http.Redirect(w, r, "/", http.StatusFound)
}

func AddTodos(w http.ResponseWriter, r *http.Request, todos *Todos, dataPath string, todo *Todo) {
	todos.Todos = append(todos.Todos, todo)

	data, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	os.WriteFile(dataPath, data, 0600)

	http.Redirect(w, r, "/", http.StatusFound)
}
