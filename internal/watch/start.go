package watch

import (
	"log"
	"time"

	"github.com/aceberg/LightAlert/internal/models"
)

// Start watching hosts
func Start(hostsMap map[string]models.Host, hashChan chan string, quit chan bool) {

	// Endless cycle with timeout
	for {
		select {
		case <-quit:
			return
		case hash := <-hashChan:

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
					}
				}
			}

			time.Sleep(time.Duration(1) * time.Second) // Cycle to check if quit
		}
	}
}
