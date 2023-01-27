package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/dchest/uniuri"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/watch"
	"github.com/aceberg/LightAlert/internal/yaml"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var oneHost models.Host

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
		alerts := strings.Join(oneHost.Alerts, " ")
		oneHost.Alerts = append([]string{}, alerts)
	}

	guiData.Hosts = append(guiData.Hosts, oneHost)

	execTemplate(w, "edit", guiData)
}

func saveHostHandler(w http.ResponseWriter, r *http.Request) {
	var oneHost models.Host

	oneHost.Name = r.FormValue("name")
	oneHost.Hash = r.FormValue("hash")
	oneHost.Interval = r.FormValue("interval")
	alerts := r.FormValue("alerts")

	if oneHost.Interval == "" {
		oneHost.Interval = "1d"
	}
	oneHost.IntSec = check.TimeToSec(oneHost.Interval)

	alertsSlice := strings.Split(alerts, " ")

	for _, a := range alertsSlice {
		if a != "" {
			oneHost.Alerts = append(oneHost.Alerts, a)
		}
	}

	log.Println("INFO: saving host", oneHost)

	HostsMap[oneHost.Hash] = oneHost
	AllHosts = check.ToStruct(HostsMap)

	close(AppConfig.Quit)

	yaml.Write(AppConfig.YamlPath, AllHosts)

	AppConfig.Quit = make(chan bool)
	go watch.Start(HostsMap, AppConfig)

	http.Redirect(w, r, "/", 302)
}

func delHostHandler(w http.ResponseWriter, r *http.Request) {

	hash := r.FormValue("hash")

	delete(HostsMap, hash)
	AllHosts = check.ToStruct(HostsMap)

	close(AppConfig.Quit)

	yaml.Write(AppConfig.YamlPath, AllHosts)

	AppConfig.Quit = make(chan bool)
	go watch.Start(HostsMap, AppConfig)

	http.Redirect(w, r, "/", 302)
}
