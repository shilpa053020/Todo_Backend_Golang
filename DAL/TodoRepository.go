package DAL

import (
	"Todo_Backend_Golang/Model"
	"context"
)

type TodoRepo interface {
	Create(ctx context.Context, todo Model.Todo) (*Model.Todo, error)
	Get(ctx context.Context)([]Model.Todo,error)
	Update(ctx context.Context ,Id string , updatedTodo Model.Todo)(string,error)
	Delete(ctx context.Context,Id string)(string,error)
}