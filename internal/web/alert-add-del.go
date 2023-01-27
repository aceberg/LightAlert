package web

import (
	"net/http"

	"github.com/aceberg/LightAlert/internal/conf"
	"github.com/aceberg/LightAlert/internal/watch"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/notify"
)

func addAlertHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	url := r.FormValue("url")

	if name != "" && url != "" {

		AppConfig.AlertMap[name] = url

		close(AppConfig.Quit)
		conf.Write(AppConfig)

		AppConfig.Quit = make(chan bool)
		go watch.Start(HostsMap, AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delAlertHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	if name != "" {

		delete(AppConfig.AlertMap, name)

		close(AppConfig.Quit)
		conf.Write(AppConfig)

		AppConfig.Quit = make(chan bool)
		go watch.Start(HostsMap, AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func testAlertHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name != "" {
		var host models.Host

		host.Name = "Test Alert"
		host.Alerts = append([]string{}, name)

		notify.Shoutrrr(host, AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}