package data

import "github.com/arora-aditya/monorepo/application-server/graph/model"

var (
	good_dependencies []*model.Dependency = []*model.Dependency{
		{
			ID:      "1",
			Name:    "b2sdk",
			Version: "1.2.3",
		},
		{
			ID:      "2",
			Name:    "django",
			Version: "4.4.3",
		},
		{
			ID:      "3",
			Name:    "urllib3",
			Version: "1.26.5",
		},
	}

	vulnerable_dependencies []*model.Dependency = []*model.Dependency{
		{
			ID:      "1",
			Name:    "b2sdk",
			Version: "1.13.1",
		},
		{
			ID:      "2",
			Name:    "requests",
			Version: "2.20.0",
		},
		{
			ID:      "3",
			Name:    "urllib3",
			Version: "1.24.3",
		},
	}

	vulnerable_devices []*model.Device = []*model.Device{
		{
			ID:           "1",
			Name:         "raspi001",
			Dependencies: vulnerable_dependencies,
		},
		{
			ID:           "2",
			Name:         "raspi002",
			Dependencies: vulnerable_dependencies,
		},
		{
			ID:           "3",
			Name:         "raspi003",
			Dependencies: vulnerable_dependencies,
		},
		{
			ID:           "5",
			Name:         "raspi005",
			Dependencies: vulnerable_dependencies,
		},
		{
			ID:           "6",
			Name:         "raspi006",
			Dependencies: vulnerable_dependencies,
		},
	}

	good_devices []*model.Device = []*model.Device{
		{
			ID:           "7",
			Name:         "raspi007",
			Dependencies: good_dependencies,
		},
		{
			ID:           "8",
			Name:         "ecserv1",
			Dependencies: good_dependencies,
		},
		{
			ID:           "9",
			Name:         "ecserv2",
			Dependencies: good_dependencies,
		},
		{
			ID:           "10",
			Name:         "ecserv3",
			Dependencies: good_dependencies,
		},
		{
			ID:           "11",
			Name:         "FPGA-21",
			Dependencies: good_dependencies,
		},
	}

	vulnerabilities []*model.Vulnerability = []*model.Vulnerability{
		{
			ID:                 "28",
			Permalink:          "https://github.com/advisories/GHSA-p867-fxfr-ph2w",
			Severity:           "MODERATE",
			Summary:            "b2-sdk-python TOCTOU application key disclosure",
			PatchAvailable:     false,
			KeyIsPatched:       false,
			Dependency:         vulnerable_dependencies[0],
			Name:               vulnerable_dependencies[0].Name,
			PatchedVersions:    []string{"1.14.1"},
			UnaffectedVersions: []string{},
			AffectedVersions:   []string{"< 1.14.1"},
			DevicesAffected:    vulnerable_devices,
		},
		{
			ID:                 "954",
			Permalink:          "https://github.com/advisories/GHSA-q2q7-5pp4-w6pg",
			Severity:           "HIGH",
			Summary:            "Catastrophic backtracking in URL authority parser when passed URL containing many @ characters",
			PatchAvailable:     true,
			KeyIsPatched:       false,
			Dependency:         vulnerable_dependencies[2],
			Name:               vulnerable_dependencies[2].Name,
			PatchedVersions:    []string{"1.26.5"},
			UnaffectedVersions: []string{},
			AffectedVersions:   []string{"< 1.26.5"},
			DevicesAffected:    vulnerable_devices,
		},
	}

	// This is the parsed JSON file that we fetch from dependabot
	all_vulnerabilities []*model.Vulnerability = []*model.Vulnerability{}
)
