package web

import (
	"net/http"

	"github.com/aceberg/LightAlert/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = AllHosts

	guiData.Themes = append([]string{}, OkIcon)
	guiData.Themes = append(guiData.Themes, ErrIcon)

	execTemplate(w, "index", guiData)
}
