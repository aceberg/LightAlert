package check

import (
	"github.com/aceberg/LightAlert/internal/models"
)

// ToMap converts struct to map
func ToMap(allHosts []models.Host) map[string]models.Host {

	m := make(map[string]models.Host)

	for _, host := range allHosts {

		host.IntSec = TimeToSec(host.Interval)

		host.Active = false

		m[host.Hash] = host
	}

	return m
}

// ToStruct converts map to struct
func ToStruct(hostsMap map[string]models.Host) []models.Host {
	var allHosts []models.Host

	for _, host := range hostsMap {
		allHosts = append(allHosts, host)
	}

	return allHosts
}
