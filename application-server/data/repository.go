package data

import (
	"context"	
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

// Repository defines the query patterns for accessing application data
type Repository interface {
	GetVulnerability(ctx context.Context, id string) (*model.Vulnerability, error)
	GetVulnerabilities(ctx context.Context, limit int, offset int) ([]*model.Vulnerability, error)

	GetDependency(ctx context.Context, name string) (*model.Dependency, error)
	GetDependencies(ctx context.Context, limit int, offset int) ([]*model.Dependency, error)

	GetDevice(ctx context.Context, id string) (*model.Device, error)
	GetDevices(ctx context.Context, limit int, offset int) ([]*model.Device, error)

	Login(input model.Login) (*model.Token, error)
	CreateUser(input model.User) (*model.Token, error)
}
