package check

import (
	"strconv"

	"github.com/aceberg/LightAlert/internal/models"
)

// ToMap converts struct to map
func ToMap(allHosts []models.Host) map[string]models.Host {

	m := make(map[string]models.Host)

	for _, host := range allHosts {

		tSec, err := TimeToSec(host.Interval)
		IfError(err)

		host.IntSec, _ = strconv.Atoi(tSec)
		host.Active = false

		m[host.Hash] = host
	}

	return m
}
