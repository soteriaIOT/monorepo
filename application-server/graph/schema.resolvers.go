package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/arora-aditya/monorepo/application-server/graph/generated"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func (r *queryResolver) Vulnerabilities(ctx context.Context) ([]*model.Vulnerability, error) {
	return r.Resolver.Vulnerabilities, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	return r.Resolver.Devices, nil
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	return r.Resolver.Images, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
