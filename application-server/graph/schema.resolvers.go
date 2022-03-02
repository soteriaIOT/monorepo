package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/arora-aditya/monorepo/application-server/graph/generated"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func (r *queryResolver) Vulnerability(ctx context.Context, id string) (*model.Vulnerability, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Vulnerabilities(ctx context.Context, limit *int, offset *int) ([]*model.Vulnerability, error) {
	return r.Resolver.Vulnerabilities, nil
}

func (r *queryResolver) Dependency(ctx context.Context, id string) (*model.Dependency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dependencies(ctx context.Context, limit *int, offset *int) ([]*model.Dependency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Device(ctx context.Context, id string) (*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Devices(ctx context.Context, limit *int, offset *int) ([]*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
