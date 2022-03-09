package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/arora-aditya/monorepo/application-server/graph/generated"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.User) (*model.Token, error) {
	return r.Repository.CreateUser(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Token, error) {
	return r.Repository.Login(input)
}

func (r *mutationResolver) UpdateVulnerabilities(ctx context.Context, input []string) ([]*model.Vulnerability, error) {
	return r.Repository.UpdateVulnerabilities(ctx, input)
}

func (r *mutationResolver) CheckVulnerabilities(ctx context.Context, input []string) ([]*model.Vulnerability, error) {
	// TODO: Send back dummy bool, we don't have any data to check at the moment
	// This is just for show
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Vulnerability(ctx context.Context, id string) (*model.Vulnerability, error) {
	return r.Repository.GetVulnerability(ctx, id)
}

func (r *queryResolver) Vulnerabilities(ctx context.Context, limit int, offset int) ([]*model.Vulnerability, error) {
	return r.Repository.GetVulnerabilities(ctx, limit, offset)
}

func (r *queryResolver) Device(ctx context.Context, id string) (*model.Device, error) {
	return r.Repository.GetDevice(ctx, id)
}

func (r *queryResolver) Devices(ctx context.Context, limit int, offset int) ([]*model.Device, error) {
	return r.Repository.GetDevices(ctx, limit, offset)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Dependency(ctx context.Context, id string) (*model.Dependency, error) {
	return r.Repository.GetDependency(ctx, id)
}
func (r *queryResolver) Dependencies(ctx context.Context, limit int, offset int) ([]*model.Dependency, error) {
	return r.Repository.GetDependencies(ctx, limit, offset)
}
