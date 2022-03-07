package data

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"strconv"
	"strings"

	"github.com/arora-aditya/monorepo/application-server/auth"
	"github.com/arora-aditya/monorepo/application-server/kafka_utils"
	"github.com/arora-aditya/monorepo/application-server/graph/model"
	"github.com/segmentio/kafka-go"
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

func (r *demoDataRepository) Login(input model.Login) (*model.Token, error) {
	scv := auth.NewDynamoSvc()
	return scv.VerifyByUsernameAndPassword(input.Username, input.Password)
}

func (r *demoDataRepository) CreateUser(input model.User) (*model.Token, error) {
	scv := auth.NewDynamoSvc()
	return scv.CreateUser(input.Name, input.Username, input.Password)
}

func (r *demoDataRepository) GetVulnerability(ctx context.Context, name string) (*model.Vulnerability, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return nil, fmt.Errorf("access denied")
	}
	for _, v := range r.Vulnerabilities {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) UpdateVulnerabilities(ctx context.Context, ids []string) ([]*model.Vulnerability, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return nil, fmt.Errorf("access denied")
	}
	for _, v := range r.Vulnerabilities {
		for _, id := range ids {
			if v.ID == id {
				for _, d := range v.DevicesAffected {
					kafka_utils.PushMessage(ctx, d.Name, v.Dependency.Name + "==" + v.PatchedVersions[0])
				}	
			}
		}
	}
	return nil, nil
}


func (r *demoDataRepository) GetVulnerabilities(ctx context.Context, limit int, offset int) ([]*model.Vulnerability, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return []*model.Vulnerability{}, fmt.Errorf("access denied")
	}
	if offset > len(r.Vulnerabilities) {
		return []*model.Vulnerability{}, nil
	}
	bound := minInt(offset+limit, len(r.Vulnerabilities))
	return r.Vulnerabilities[offset:bound], nil
}

func (r *demoDataRepository) GetDependency(ctx context.Context, id string) (*model.Dependency, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return nil, fmt.Errorf("access denied")
	}
	for _, d := range r.Dependencies {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) GetDependencies(ctx context.Context, limit int, offset int) ([]*model.Dependency, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return []*model.Dependency{}, fmt.Errorf("access denied")
	}
	if offset > len(r.Dependencies) {
		return []*model.Dependency{}, nil
	}
	bound := minInt(offset+limit, len(r.Dependencies))
	return r.Dependencies[offset:bound], nil
}

func (r *demoDataRepository) GetDevice(ctx context.Context, id string) (*model.Device, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return nil, fmt.Errorf("access denied")
	}
	for _, d := range r.Devices {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, nil
}

func (r *demoDataRepository) GetDevices(ctx context.Context, limit int, offset int) ([]*model.Device, error) {
	user := auth.GetAuthFromContext(ctx)
	if user.Username == "" {
		return []*model.Device{}, fmt.Errorf("access denied")
	}
	if offset > len(r.Devices) {
		return []*model.Device{}, nil
	}
	bound := minInt(offset+limit, len(r.Devices))
	return r.Devices[offset:bound], nil
}

func (r *demoDataRepository) UpdateDeviceDependencies(ctx context.Context, device_name string, dependencies string) error {
	done := false
	for _, d := range r.Devices {
		if d.Name == device_name {
			d.Dependencies, d.Vulnerabilities = r.ParseDepencencies(ctx, dependencies)
			for _, v := range d.Vulnerabilities {
				r.addDeviceToVulnerability(v, d)
			}
			done = true
		}
	}
	max_id := 0
	for _, d := range r.Devices {
		num, _ := strconv.Atoi(d.ID)
		max_id = maxInt(num, max_id)
	}
	if !done {
		dependencies, vulnerabilities := r.ParseDepencencies(ctx, dependencies)
		new_device := &model.Device{
			ID: strconv.Itoa(max_id + 1),
			Name: device_name,
			Dependencies: dependencies,
			Vulnerabilities: vulnerabilities,
		}
		r.Devices = append(r.Devices, new_device)
		for _, v := range new_device.Vulnerabilities {
			r.addDeviceToVulnerability(v, new_device)
		}
		
	}
	
	return nil
}

func (r *demoDataRepository) ParseDepencencies(ctx context.Context, dependencies string) ([]*model.Dependency, []*model.Vulnerability) {
	dependencies_as_list := strings.Split(dependencies, "\n")
	deps := []*model.Dependency{};
	vulnerabilities := []*model.Vulnerability{};
	max_id := 0
	for _, d := range r.Dependencies {
		num, _ := strconv.Atoi(d.ID)
		max_id = maxInt(num, max_id)
	}
	for _, dependency_string := range dependencies_as_list {
		parsed := strings.Split(dependency_string, "==")
		dependency_name := ""
		version := ""
		if len(parsed) != 2 {
			continue
		} else {
			dependency_name = parsed[0]
			version = parsed[1]
		}
		done := false
		for _, d := range r.Dependencies {
			if dependency_name == d.Name && version == d.Version {
				deps = append(deps, d)
				isVulnerable, vuln := r.isVulnerable(d)
				if isVulnerable {
					vulnerabilities = append(vulnerabilities, vuln)
				}
				done = true
			}
		}
		if !done {
			new_dep := &model.Dependency{
				ID: strconv.Itoa(max_id + 1),
				Name:     dependency_name,
				Version:  version,
			}
			deps = append(deps, new_dep)
			isVulnerable, vuln := r.isVulnerable(new_dep)
			if isVulnerable {
				vulnerabilities = append(vulnerabilities, vuln)
			}
			r.Dependencies = append(r.Dependencies, new_dep)
			max_id++
		}
	}
	return deps, vulnerabilities
}

func (r *demoDataRepository) ReadMessage(ctx context.Context, wg *sync.WaitGroup) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_IP")},
		// No groupID because we want to parse all the messages from the topic
		// and come back to current state on every restart since our memory gets restart on reboot
		// GroupID:   "application-server",
		Topic:     "device-requirements",
		MinBytes:  10e2, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	for {
		select {
			case <-ctx.Done():
				log.Println("Closing Kafka Reader")
				if err := reader.Close(); err != nil {
					log.Println("Failed to close reader:", err)
				}
				wg.Done()
				return
			default:
				// The same context needs to be passed so that we can terminate on Ctrl C gracefully
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					log.Println("Closing Kafka Reader")
					if err := reader.Close(); err != nil {
						log.Println("Failed to close reader:", err)
					}
					wg.Done()
					return
				}
				r.UpdateDeviceDependencies(ctx, string(m.Key), string(m.Value))
		}
	}
}

func (r *demoDataRepository) isVulnerable(dependency *model.Dependency) (bool, *model.Vulnerability) {
	for _, vulnerability := range r.Vulnerabilities {
		if vulnerability.Dependency.Name == dependency.Name && dependency.Version <= vulnerability.PatchedVersions[0] {
			return true, vulnerability
		}
	}
	return false, nil
}

func (r *demoDataRepository) addDeviceToVulnerability(vulnerability *model.Vulnerability, device *model.Device) {
	exists := false
	for _, d := range vulnerability.DevicesAffected {
		if d.Name == device.Name {
			exists = true
			return
		}
	}
	if !exists {
		vulnerability.DevicesAffected = append(vulnerability.DevicesAffected, device)
	}
}


func maxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
