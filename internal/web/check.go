package web

import (
	"log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/check"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {

	hash := r.URL.Query().Get("hash")

	_, exist := HostsMap[hash]
	if exist {
		HashChan <- hash

		ipAddress := readUserIP(r)
		userAgent := r.UserAgent()

		log.Println("FROM IP:", ipAddress)
		log.Println("USER-AGENT:", userAgent)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		check.IfError(err)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func readUserIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}
