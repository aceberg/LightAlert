package watch

import (
	"log"
	"time"

	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/notify"
)

// Start watching hosts
func Start(hostsMap map[string]models.Host, conf models.Conf) {

	log.Println("INFO: watching hosts", hostsMap)

	// Endless cycle with timeout
	for {
		select {
		case <-conf.Quit:
			return
		case hash := <-conf.HashChan:

			host := hostsMap[hash]
			host.LastSeen = time.Now()
			if !host.Active {
				notify.Up(host, conf.AlertMap)
				host.Active = true
			}
			log.Println("INFO: check from", host)
			// Needs pause for host.LastSeen not be empty. WHY???
			time.Sleep(time.Duration(1) * time.Second)
			hostsMap[hash] = host

		default:
			nowDate := time.Now()

			for hash, host := range hostsMap {

				// log.Println("CHECKING:", host)
				if host.Active {
					lastDate := host.LastSeen
					plusDate := lastDate.Add(time.Duration(host.IntSec) * time.Second)

					if nowDate.After(plusDate) {
						conf.OffChan <- hash
						host.Active = false
						hostsMap[hash] = host

						log.Println("ALERT:", host)

						var rec models.Record
						rec.Date = time.Now().Format("2006-01-02 15:04:05")
						rec.Name = host.Name
						rec.Hash = host.Hash
						rec.State = "alert"
						db.Insert(conf.DB, rec)

						notify.Down(host, conf.AlertMap)
					}
				}
			}

			time.Sleep(time.Duration(1) * time.Second) // Timeout
		}
	}
}
