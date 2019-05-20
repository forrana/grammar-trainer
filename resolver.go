package api

import (
	"context"

	"github.com/forrana/gramma-trainer/api/data"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Exercise() ExerciseResolver {
	return &exerciseResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type exerciseResolver struct{ *Resolver }

func (r *exerciseResolver) User(ctx context.Context, obj *data.Exercise) (*User, error) {
	panic("not implemented")
}
func (r *exerciseResolver) Tags(ctx context.Context, obj *data.Exercise) ([]Tag, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateExercise(ctx context.Context, input NewExercise) (*data.Exercise, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateSection(ctx context.Context, input NewSection) (*Section, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTag(ctx context.Context, input NewTag) (*Tag, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTrial(ctx context.Context, input NewTrial) (*Trial, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteExercise(ctx context.Context, ids []int) ([]*data.Exercise, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteSections(ctx context.Context, ids []int) ([]*Section, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTags(ctx context.Context, ids []int) ([]*Tag, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTrials(ctx context.Context, ids []int) ([]*Trial, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Exercises(ctx context.Context) ([]data.Exercise, error) {
	panic("not implemented")
}
func (r *queryResolver) Sections(ctx context.Context) ([]Section, error) {
	panic("not implemented")
}
func (r *queryResolver) Tags(ctx context.Context) ([]Tag, error) {
	panic("not implemented")
}
