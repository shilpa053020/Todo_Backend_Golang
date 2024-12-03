# Todo_Backend_Golang
## A simple backend application built using Golang to perform basic CRUD (Create, Read, Update, Delete) operations for a Todo application.

## Features
  1.Create Todos: Add new tasks to your todo list.
  2.Read Todos: Fetch all existing tasks.
  3.Update Todos: Modify existing tasks by their ID.  
  4.Delete Todos: Remove tasks from the list.

## Technologies Used
  1.Golang: Backend logic
  2.MongoDB: Database for storing todos
  3.Context API: For managing timeout and cancellation in requests
  4.Gorilla Mux: HTTP request routing

## Getting Started
# Follow these steps to set up and run the project locally.

## Make sure you have the following installed on your system:
  1.Go (v1.18 or later)
  2.MongoDB (Running on localhost)
  3.Git


## Clone the repository:
  1.git clone https://github.com/yourusername/Todo_Backend_Golang.git
  2.cd Todo_Backend_Golang

## Install dependencies:
  go mod tidy

## Environment Variables:
#Create a .env file in the root directory with the following content:
  MONGO_URI=mongodb://localhost:27017
  PORT = 5000

## Run the application:
  go run main.go
