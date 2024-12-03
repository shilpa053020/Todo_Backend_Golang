package Service

import (
	"Todo_Backend_Golang/DAL"
	"Todo_Backend_Golang/Model"
	utils "Todo_Backend_Golang/Utils"
	"time"
)

type TodoService struct {
	todoRepo DAL.TodoRepo
}

func NewTodoService(Todorepo DAL.TodoRepo) *TodoService{
	return &TodoService{todoRepo: Todorepo}
}

//type-name-params-return
func (service TodoService)CreateTodo(todo Model.Todo) (*Model.Todo,error)  {
	ctx,cancel:= utils.GetContext(5*time.Second)
	defer cancel()
	return service.todoRepo.Create(ctx,todo)
	
}

func (service TodoService)GetAllTodos()([]Model.Todo,error)  {
	ctx,cancel:= utils.GetContext(5*time.Second)
	defer cancel()
    return service.todoRepo.Get(ctx)
}

func(service TodoService)UpdateTodo(id string,updatedTodo Model.Todo)(string,error){
	ctx,cancel:= utils.GetContext(5*time.Second)
	defer cancel()
	return service.todoRepo.Update(ctx,id,updatedTodo)
}

func(service TodoService)DeleteTodo(id string)(string,error){
	ctx,cancel:= utils.GetContext(5*time.Second)
	defer cancel()
	return service.todoRepo.Delete(ctx,id)
}