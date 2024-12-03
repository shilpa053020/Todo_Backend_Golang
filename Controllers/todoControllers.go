package Controllers

import (
	"Todo_Backend_Golang/Model"
	Service "Todo_Backend_Golang/Services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoControllers struct {
	todoService *Service.TodoService
}
func NewTodoHandler(services *Service.TodoService) *TodoControllers{
	 fmt.Println("Controllers")
	 return &TodoControllers{todoService: services}
}

func(tc *TodoControllers) CreateATodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create")
	var Todo Model.Todo
    err:= json.NewDecoder(r.Body).Decode(&Todo)
	if (err != nil) {
		http.Error(w,"Invalid Input:"+err.Error(),http.StatusBadRequest)
		return
	}
	createdTodo,err := tc.todoService.CreateTodo(Todo)
	if(err != nil){
		http.Error(w,"Failed to create:" +err.Error(),http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(createdTodo) 
}

func (tc *TodoControllers) GetAllTodos(w http.ResponseWriter, r *http.Request){
	todos , err := tc.todoService.GetAllTodos()
	if(err != nil){
		http.Error(w,"Failed to get:" +err.Error(),http.StatusBadRequest)
		return
	}
    json.NewEncoder(w).Encode(todos)
}

func (tc *TodoControllers) UpdateATodo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTodo Model.Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	result , err := tc.todoService.UpdateTodo(id,updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func(tc *TodoControllers)DeleteTodo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	result,err:= tc.todoService.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}