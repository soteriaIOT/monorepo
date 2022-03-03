package data

import "github.com/arora-aditya/monorepo/application-server/graph/model"

// Repository defines the query patterns for accessing application data
type Repository interface {
	GetVulnerability(id string) (*model.Vulnerability, error)
	GetVulnerabilities(limit int, offset int) ([]*model.Vulnerability, error)

	GetDependency(id string) (*model.Dependency, error)
	GetDependencies(limit int, offset int) ([]*model.Dependency, error)

	GetDevice(id string) (*model.Device, error)
	GetDevices(limit int, offset int) ([]*model.Device, error)
}
