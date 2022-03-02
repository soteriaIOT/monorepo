package main

import "github.com/arora-aditya/monorepo/application-server/graph/model"

// DataRepository defines the query patterns for accessing application data
type DataRepository interface {
	GetVulnerability(id string) (*model.Vulnerability, error)
	GetVulnerabilities(limit int, offset int) []*model.Vulnerability

	GetDependency(id string) (*model.Dependency, error)
	GetDependencies(limit int, offset int) []*model.Dependency

	GetDevice(id string) (*model.Device, error)
	GetDevices(limit int, offset int) []*model.Device
}
