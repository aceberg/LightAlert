package web

import (
	// "log"
	"net/http"
	"strings"

	"github.com/dchest/uniuri"

	// "github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var oneHost models.Host
	var recNext models.Record

	guiData.Config = AppConfig

	hash := r.URL.Query().Get("hash")

	if hash == "new" {
		hash = uniuri.New()

		_, exist := HostsMap[hash]
		if exist {
			hash = uniuri.New()
		}
		oneHost.Hash = hash

	} else {
		oneHost = HostsMap[hash]
		alerts := strings.Join(oneHost.Alerts, " ")
		oneHost.Alerts = append([]string{}, alerts)

		recNext = models.Record{}
		for _, rec := range LogRecords {
			if rec.Hash == hash && rec.State == "ok" {
				recNext = rec
			}
			if rec.Hash == hash && rec.State == "alert" {
				if recNext != (models.Record{}) {
					guiData.Records = append(guiData.Records, recNext)
				}
				guiData.Records = append(guiData.Records, rec)
				recNext = models.Record{}
			}
		}
	}

	guiData.Hosts = append(guiData.Hosts, oneHost)

	execTemplate(w, "edit", guiData)
}
