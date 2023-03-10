package web

import (
	"log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/conf"
	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/watch"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.DB = r.FormValue("db")
	AppConfig.DelOld = r.FormValue("delold")
	AppConfig.HName = r.FormValue("hname")
	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Show = r.FormValue("show")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.YamlPath = r.FormValue("yamlpath")

	AppConfig.Quit <- true
	conf.Write(AppConfig)
	go watch.Start(HostsMap, AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: deleting all records from DB")

	db.Clear(AppConfig.DB)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
