package check

import (
	"github.com/aceberg/LightAlert/internal/models"
)

// ToMap converts struct to map
func ToMap(allHosts []models.Host) map[string]models.Host {

	m := make(map[string]models.Host)

	for _, oneHost := range allHosts {
		m[oneHost.Hash] = oneHost
	}

	return m
}
