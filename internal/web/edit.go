package web

import (
	// "log"
	"net/http"

	"github.com/dchest/uniuri"

	// "github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var oneHost models.Host
	var recNext models.Record
	var alCheck models.EditAlert

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

	for name := range AppConfig.AlertMap {
		alCheck.Name = name
		alCheck.Check = false

		for _, alert := range oneHost.Alerts {
			if alert == name {
				alCheck.Check = true
			}
		}

		guiData.AlertCheck = append(guiData.AlertCheck, alCheck)
	}

	guiData.Hosts = append(guiData.Hosts, oneHost)

	execTemplate(w, "edit", guiData)
}
