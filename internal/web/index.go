package web

import (
	"net/http"
	"sort"

	"github.com/aceberg/LightAlert/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	sort.Slice(AllHosts, func(i, j int) bool {
		return AllHosts[i].Name < AllHosts[j].Name
	})

	guiData.Hosts = AllHosts

	execTemplate(w, "index", guiData)
}
