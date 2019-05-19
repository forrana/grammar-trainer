package gramma_trainer

import (
	"context"
    "fmt"
    "math/rand"
	"github.com/forrana/gramma-trainer/api/data"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	todos []data.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*data.Todo, error) {
	todo := &data.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, *todo)
	return todo, data.DB.Create(&todo).Error
}
func (r *mutationResolver) DeleteTodos(ctx context.Context, ids []int) ([]*data.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ExercisesList(ctx context.Context) (todos []data.Todo, err error) {
	return todos, data.DB.Find(&todos).Error
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *data.Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
