package web

import (
	"net/http"

	"github.com/aceberg/LightAlert/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	execTemplate(w, "index", guiData)
}