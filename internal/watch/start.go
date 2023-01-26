package watch

import (
	"log"
	"time"

	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/notify"
)

// Start watching hosts
func Start(hostsMap map[string]models.Host, conf models.Conf) {

	// Endless cycle with timeout
	for {
		select {
		case <-conf.Quit:
			return
		case hash := <-conf.HashChan:

			host := hostsMap[hash]
			host.LastSeen = time.Now()
			host.Active = true
			hostsMap[hash] = host

		default:
			nowDate := time.Now()

			for hash, host := range hostsMap {
				if host.Active {
					lastDate := host.LastSeen
					plusDate := lastDate.Add(time.Duration(host.IntSec) * time.Second)

					if nowDate.After(plusDate) {
						host.Active = false
						hostsMap[hash] = host

						log.Println("ALERT:", host)
						notify.Shoutrrr(host)
					}
				}
			}

			time.Sleep(time.Duration(1) * time.Second) // Cycle to check if quit
		}
	}
}
