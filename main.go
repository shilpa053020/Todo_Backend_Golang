package main

import (
	"Todo_Backend_Golang/Configure"
	"Todo_Backend_Golang/Controllers"
	"Todo_Backend_Golang/DAL"
	"Todo_Backend_Golang/Routes"
	Service "Todo_Backend_Golang/Services"
	"fmt"
	"net/http"
    "log"
	"os"
	"github.com/gorilla/mux"
);

func main() {
	//Connection to db
	configure.ConnectToDb();

	//instances
	todoDal := DAL.NewMongoDbDal(configure.Client,"Todo_app","Todos")
	todoService := Service.NewTodoService(todoDal)
	todoController := Controllers.NewTodoHandler(todoService)
    
	//Setup port 
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080" 
	}

    //routes setup
	router := mux.NewRouter()
	Routes.SetUpTodoRoutes(router,todoController)

	//server connection
	fmt.Printf("Server is listening on port 5001...")
	err := http.ListenAndServe(":"+Port, router)
	if err != nil {
		log.Fatal("Server failed to start:", err)
    }
	
}

	