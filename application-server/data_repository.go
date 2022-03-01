package main

import "github.com/arora-aditya/monorepo/application-server/graph/model"

// DataRepository defines the query patterns for accessing application data
type DataRepository interface {
	GetVulnerability(id string) (*model.Vulnerability, error)
	FindVulnerabilities(idPrefix string, onlySeverities []model.Severity, ordering model.Ordering) []*model.Vulnerability

	GetDependency(id string) (*model.Dependency, error)
	FindDependencies(namePrefix string, ordering model.Ordering) []*model.Dependency

	GetImage(id string) (*model.Image, error)
	FindImages(repository string, ordering model.Ordering) []*model.Image

	GetDevice(id string) (*model.Device, error)
	FindDevices(namePrefix string, ordering model.Ordering) []*model.Device
}
