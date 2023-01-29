package web

import (
	"net/http"

	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Records = db.Select(AppConfig.DB)

	guiData.Themes = append([]string{}, OkIcon)
	guiData.Themes = append(guiData.Themes, ErrIcon)

	execTemplate(w, "log", guiData)
}
