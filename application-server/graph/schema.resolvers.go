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

func (r *queryResolver) Vulnerabilities(ctx context.Context, idPrefix *string, onlySeverities []model.Severity, ordering *model.Ordering) ([]*model.Vulnerability, error) {
	return r.Resolver.Vulnerabilities, nil
}

func (r *queryResolver) Device(ctx context.Context, id string) (*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Devices(ctx context.Context, namePrefix *string, ordering *model.Ordering) ([]*model.Device, error) {
	return r.Resolver.Devices, nil
}

func (r *queryResolver) Image(ctx context.Context, id string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Images(ctx context.Context, namePrefix *string, ordering *model.Ordering) ([]*model.Image, error) {
	return r.Resolver.Images, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
