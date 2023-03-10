package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var filteredRecords []models.Record

	name := r.URL.Query().Get("name")
	state := r.URL.Query().Get("state")
	showStr := r.URL.Query().Get("show")

	if name == "" && state == "" {
		filteredRecords = LogRecords
	} else {
		for _, rec := range LogRecords {
			if (state == "" && rec.Name == name) || (name == "" && rec.State == state) || (rec.Name == name && rec.State == state) {
				filteredRecords = append(filteredRecords, rec)
			}
		}
	}

	guiData.Config = AppConfig
	guiData.Len = len(filteredRecords)

	if showStr != "" {
		AppConfig.Show = showStr
	}

	show, err := strconv.Atoi(AppConfig.Show)

	if show > guiData.Len || check.IfError(err) {
		show = guiData.Len
	}

	guiData.Records = filteredRecords[0:show]

	for _, host := range AllHosts {
		guiData.Themes = append(guiData.Themes, host.Name)
	}

	execTemplate(w, "log", guiData)
}
