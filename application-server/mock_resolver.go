package main

import (
	"github.com/arora-aditya/monorepo/application-server/graph"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

func NewMockResolver() *graph.Resolver {
	dependencies := []*model.Dependency{
		{
			ID:      "1",
			Name:    "b2sdk",
			Version: "1.2.3",
		},
	}

	vulnerabilities := []*model.Vulnerability{
		{
			Permalink:              "https://github.com/advisories/GHSA-p867-fxfr-ph2w",
			Severity:               "MODERATE",
			Summary:                "b2-sdk-python TOCTOU application key disclosure",
			VulnerableVersionRange: "< 1.14.1",
			PatchAvailable:         true,
			KeyIsPatched:           false,
			ID:                     "CVE-2022-23651",
			Dependencies:           dependencies,
			PatchedVersions:        []string{"1.14.1"},
			UnaffectedVersions:     []string{},
			AffectedVersions:       []string{"< 1.14.1"},
		},
	}

	devices := []*model.Device{
		{
			ID:   "1",
			Name: "acme-security-camera",
		},
	}

	return &graph.Resolver{
		Dependencies:    dependencies,
		Vulnerabilities: vulnerabilities,
		Devices:         devices,
	}
}
