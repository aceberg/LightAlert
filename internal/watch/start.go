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
			// log.Println("TIME:", host.LastSeen)
			// Needs pause for host.LastSeen not be empty. WHY???
			time.Sleep(time.Duration(1) * time.Second)
			host.Active = true
			hostsMap[hash] = host

			// log.Println("NOW ACTIVE:", hostsMap)
		default:
			time.Sleep(time.Duration(1) * time.Second)

			nowDate := time.Now()

			for hash, host := range hostsMap {
				if host.Active {
					// log.Println("CHECK:", hostsMap)
					lastDate := host.LastSeen
					plusDate := lastDate.Add(time.Duration(host.IntSec) * time.Second)

					if nowDate.After(plusDate) {
						conf.OffChan <- hash
						host.Active = false
						hostsMap[hash] = host

						log.Println("ALERT:", host)
						notify.Shoutrrr(host, conf)
					}
				}
			}
		}
	}
}
