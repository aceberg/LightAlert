package web

import (
	"net/http"
	"time"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {

	hash := r.URL.Query().Get("hash")

	host, exist := HostsMap[hash]
	if exist {
		AppConfig.HashChan <- hash

		host.Active = true
		HostsMap[hash] = host
		AllHosts = check.ToStruct(HostsMap)

		var rec models.Record
		rec.Date = time.Now().Format("2006-01-02 15:04:05")
		rec.Name = host.Name
		rec.Hash = host.Hash
		rec.IP = readUserIP(r)
		rec.Agent = r.UserAgent()

		db.Insert(AppConfig.DB, rec)

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
