package gramma_trainer

import (
	"context"
    "fmt"
    "math/rand"
	"github.com/forrana/gramma-trainer/api/data"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	exercises []data.Exercise
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Exercise() ExerciseResolver {
	return &exerciseResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateExercise(ctx context.Context, input NewExercise) (*data.Exercise, error) {
	exercise := &data.Exercise{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.exercises = append(r.exercises, *exercise)
	return exercise, data.DB.Create(&exercise).Error
}
func (r *mutationResolver) DeleteExercise(ctx context.Context, ids []int) ([]*data.Exercise, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ExercisesList(ctx context.Context) (exercises []data.Exercise, err error) {
	return exercises, data.DB.Find(&exercises).Error
}

type exerciseResolver struct{ *Resolver }

func (r *exerciseResolver) User(ctx context.Context, obj *data.Exercise) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
