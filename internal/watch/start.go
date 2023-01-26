package watch

import (
	"log"
	"time"

	"github.com/aceberg/LightAlert/internal/models"
)

// Start watching hosts
func Start(hostsMap map[string]models.Host, hashChan chan string, quit chan bool) {

	// hostMap := make(map[string]models.Host)

	// Endless cycle with timeout
	for {
		select {
		case <-quit:
			return
		case hash := <-hashChan:

			_, exist := hostsMap[hash]
			if exist {
				log.Println("Host in map, update LastSeen")
			}
		default:
			log.Println("CHANNEL EMPTY")

			time.Sleep(time.Duration(2) * time.Second) // Cycle to check if quit
		}
	}
}
