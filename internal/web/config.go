package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/conf"
	"github.com/aceberg/LightAlert/internal/watch"
	// "github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	// file, err := TemplHTML.ReadFile(TemplPath + "version")
	// check.IfError(err)

	// version := string(file)
	// guiData.File = version[8:]

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.DB = r.FormValue("db")
	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.YamlPath = r.FormValue("yamlpath")
	AppConfig.LogPath = r.FormValue("logpath")

	close(AppConfig.Quit)
	conf.Write(AppConfig)

	AppConfig.Quit = make(chan bool)
	go watch.Start(HostsMap, AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

// func clearHandler(w http.ResponseWriter, r *http.Request) {

// 	log.Println("INFO: deleting all plays from DB")

// 	db.Clear(AppConfig.DB)

// 	http.Redirect(w, r, r.Header.Get("Referer"), 302)
// }
