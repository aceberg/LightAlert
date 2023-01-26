package yaml

import (
	// "fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

// Read - read .yaml file to struct
func Read(path string) []models.Host {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var hosts []models.Host
	err = yaml.Unmarshal(file, &hosts)
	check.IfError(err)

	return hosts
}

// Write - write struct to  .yaml file
func Write(path string, hosts []models.Host) {
	type YamlHost struct {
		Name     string   `yaml:"name"`
		Hash     string   `yaml:"hash"`
		Interval string   `yaml:"interval"`
		Alerts   []string `yaml:"alerts"`
	}
	var yamlHosts []YamlHost
	var newHost YamlHost

	for _, oneHost := range hosts {
		newHost.Name = oneHost.Name
		newHost.Hash = oneHost.Hash
		newHost.Interval = oneHost.Interval
		newHost.Alerts = oneHost.Alerts
		yamlHosts = append(yamlHosts, newHost)
	}

	yamlData, err := yaml.Marshal(&yamlHosts)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new hosts file to", path, "\n", string(yamlData))
}
