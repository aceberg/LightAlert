package web

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var filteredRecords []models.Record

	search := r.URL.Query().Get("search")

	for _, rec := range LogRecords {
		if inString(rec.Name, search) || inString(rec.Date, search) || inString(rec.IP, search) || inString(rec.Agent, search) {
			filteredRecords = append(filteredRecords, rec)
		}
	}

	guiData.Config = AppConfig
	guiData.Len = len(filteredRecords)

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

func inString(str1 string, str2 string) bool {
	return strings.Contains(
		strings.ToLower(str1),
		strings.ToLower(str2),
	)
}
