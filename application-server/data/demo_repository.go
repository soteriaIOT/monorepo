package data

import (
	"math"

	"github.com/arora-aditya/monorepo/application-server/auth"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func NewDemoRepository() Repository {
	for _, s := range vulnerable_devices {
		s.Vulnerabilities = vulnerabilities
	}

	return &demoDataRepository{
		Dependencies:    append(vulnerable_dependencies, good_dependencies...),
		Vulnerabilities: vulnerabilities,
		Devices:         append(vulnerable_devices, good_devices...),
	}
}

type demoDataRepository struct {
	Vulnerabilities []*model.Vulnerability
	Dependencies    []*model.Dependency
	Devices         []*model.Device
}

func (r *demoDataRepository) Login(input model.Login) (bool, error) {
	scv := auth.NewDynamoSvc()
	return scv.VerifyByUsernameAndPassword(input.Username, input.Password), nil
}

func (r *demoDataRepository) GetVulnerability(name string) (*model.Vulnerability, error) {
	for _, v := range r.Vulnerabilities {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) GetVulnerabilities(limit int, offset int) ([]*model.Vulnerability, error) {
	if offset > len(r.Vulnerabilities) {
		return []*model.Vulnerability{}, nil
	}
	bound := minInt(offset+limit, len(r.Vulnerabilities))
	return r.Vulnerabilities[offset:bound], nil
}

func (r *demoDataRepository) GetDependency(id string) (*model.Dependency, error) {
	for _, d := range r.Dependencies {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) GetDependencies(limit int, offset int) ([]*model.Dependency, error) {
	if offset > len(r.Dependencies) {
		return []*model.Dependency{}, nil
	}
	bound := minInt(offset+limit, len(r.Dependencies))
	return r.Dependencies[offset:bound], nil
}

func (r *demoDataRepository) GetDevice(id string) (*model.Device, error) {
	for _, d := range r.Devices {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) GetDevices(limit int, offset int) ([]*model.Device, error) {
	if offset > len(r.Devices) {
		return []*model.Device{}, nil
	}
	bound := minInt(offset+limit, len(r.Devices))
	return r.Devices[offset:bound], nil
}

func minInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
