package web

import (
	"log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/watch"
	"github.com/aceberg/LightAlert/internal/yaml"
)

func saveHostHandler(w http.ResponseWriter, r *http.Request) {
	var oneHost models.Host

	oneHost.Name = r.FormValue("name")
	oneHost.Hash = r.FormValue("hash")
	oneHost.Interval = r.FormValue("interval")
	alerts := r.PostForm["alerts"]

	if oneHost.Name == "" {
		oneHost.Name = oneHost.Hash
	}

	if oneHost.Interval == "" {
		oneHost.Interval = "1d"
	}
	oneHost.IntSec = check.TimeToSec(oneHost.Interval)

	for _, a := range alerts {
		if a != "" {
			oneHost.Alerts = append(oneHost.Alerts, a)
		}
	}

	log.Println("INFO: saving host", oneHost)

	HostsMap[oneHost.Hash] = oneHost
	AllHosts = check.ToStruct(HostsMap)

	AppConfig.Quit <- true
	yaml.Write(AppConfig.YamlPath, AllHosts)
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
