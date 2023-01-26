package web

import (
	"log"
	"net/http"
	// "github.com/aceberg/LightAlert/internal/models"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {

	hash := r.URL.Query().Get("hash")

	log.Println("CHECK FOR HASH:", hash)

	for _, host := range AllHosts {
		if hash == host.Hash {

			log.Println("CHECK!")
			log.Println(r)
		}
	}
}
