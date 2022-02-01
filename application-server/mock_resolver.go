package main

import (
	"github.com/arora-aditya/monorepo/application-server/graph"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func NewMockResolver() *graph.Resolver {
	dependencies := []*model.Dependency{
		{
			ID:      "1",
			Name:    "requests",
			Version: "1.2.3",
		},
	}

	vulnerabilities := []*model.Vulnerability{
		{
			ID:          "1",
			Description: "foo",
			Dependency:  dependencies[0],
		},
	}

	images := []*model.Image{
		{
			ID:           "1",
			Repository:   "bar",
			Tag:          "latest",
			Dependencies: dependencies,
		},
	}

	devices := []*model.Device{
		{
			ID:    "1",
			Name:  "acme-security-camera",
			Image: images[0],
		},
	}

	return &graph.Resolver{
		Dependencies:    dependencies,
		Vulnerabilities: vulnerabilities,
		Images:          images,
		Devices:         devices,
	}
}
