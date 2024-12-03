package Routes

import (
	"Todo_Backend_Golang/Controllers"
	"github.com/gorilla/mux"
)
    
func SetUpTodoRoutes (router *mux.Router,controllers *Controllers.TodoControllers){
	router.HandleFunc("/todos", controllers.CreateATodo).Methods("POST")
	router.HandleFunc("/",controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/{id}",controllers.UpdateATodo).Methods("PUT")
	router.HandleFunc("/{id}",controllers.DeleteTodo).Methods("DELETE")
}