// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Dependency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Device struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Dependencies    []*Dependency    `json:"dependencies"`
	Vulnerabilities []*Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	ID                 int         `json:"id"`
	Permalink          string      `json:"permalink"`
	Severity           Severity    `json:"severity"`
	Summary            string      `json:"summary"`
	PatchAvailable     bool        `json:"patch_available"`
	KeyIsPatched       bool        `json:"key_is_patched"`
	Name               string      `json:"name"`
	Dependency         *Dependency `json:"dependency"`
	PatchedVersions    []string    `json:"patched_versions"`
	UnaffectedVersions []string    `json:"unaffected_versions"`
	AffectedVersions   []string    `json:"affected_versions"`
	DevicesAffected    []*Device   `json:"devices_affected"`
}

type Severity string

const (
	SeverityLow      Severity = "LOW"
	SeverityModerate Severity = "MODERATE"
	SeverityHigh     Severity = "HIGH"
)

var AllSeverity = []Severity{
	SeverityLow,
	SeverityModerate,
	SeverityHigh,
}

func (e Severity) IsValid() bool {
	switch e {
	case SeverityLow, SeverityModerate, SeverityHigh:
		return true
	}
	return false
}

func (e Severity) String() string {
	return string(e)
}

func (e *Severity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Severity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Severity", str)
	}
	return nil
}

func (e Severity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
