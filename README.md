# Todo_Backend_Golang
### A simple backend application built using Golang to perform basic CRUD (Create, Read, Update, Delete) operations for a Todo application.

## Features
  - Create Todos: Add new tasks to your todo list.
  - Read Todos: Fetch all existing tasks.
  - Update Todos: Modify existing tasks by their ID.  
  - Delete Todos: Remove tasks from the list.

## Technologies Used
  - Golang: Backend logic
  - MongoDB: Database for storing todos
  - Context API: For managing timeout and cancellation in requests
  - Gorilla Mux: HTTP request routing

## Getting Started
### Follow these steps to set up and run the project locally.

## Make sure you have the following installed on your system:
  - Go (v1.18 or later)
  - MongoDB (Running on localhost)
  - Git


## Clone the repository:
  - git clone https://github.com/yourusername/Todo_Backend_Golang.git
  - cd Todo_Backend_Golang

## Install dependencies:
  go mod tidy

## Environment Variables:
#Create a .env file in the root directory with the following content:
  MONGO_URI=mongodb://localhost:27017
  PORT = 5000

## Run the application:
  go run main.go
