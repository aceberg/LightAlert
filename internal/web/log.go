package web

import (
	"net/http"

	"github.com/aceberg/LightAlert/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Records = LogRecords

	execTemplate(w, "log", guiData)
}
