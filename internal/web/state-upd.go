package web

import (
	"time"

	"github.com/aceberg/LightAlert/internal/check"
)

func stateUpdate() {

	// Endless cycle with timeout
	for {
		select {
		case hash := <-AppConfig.OffChan:
			host := HostsMap[hash]
			host.Active = false
			HostsMap[hash] = host

			AllHosts = check.ToStruct(HostsMap)

		default:
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
