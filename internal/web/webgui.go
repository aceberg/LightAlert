package web

import (
	"log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/conf"
	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/watch"
	"github.com/aceberg/LightAlert/internal/yaml"
)

// Gui - start web server
func Gui(config models.Conf) {
	AppConfig = mergeConfig(config)

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	db.Create(AppConfig.DB)
	LogRecords = db.Select(AppConfig.DB)

	AllHosts = yaml.Read(AppConfig.YamlPath)
	HostsMap = check.ToMap(AllHosts)

	go watch.Start(HostsMap, AppConfig)
	go stateUpdate() // Update AllHosts.Active for index page

	address := AppConfig.Host + ":" + AppConfig.Port
	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/alert_add/", addAlertHandler)
	http.HandleFunc("/alert_del/", delAlertHandler)
	http.HandleFunc("/alert_test/", testAlertHandler)
	http.HandleFunc("/check/", checkHandler)
	http.HandleFunc("/clear/", clearHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/del_host/", delHostHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/save_host/", saveHostHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}

func mergeConfig(config models.Conf) models.Conf {

	newConfig := conf.Get(config.ConfPath)
	newConfig.ConfPath = config.ConfPath

	if config.YamlPath != "" {
		newConfig.YamlPath = config.YamlPath
	}
	if config.DB != "" {
		newConfig.DB = config.DB
	}

	newConfig.Icon = Icon

	newConfig.Quit = make(chan bool, 10)
	newConfig.HashChan = make(chan string, 100)
	newConfig.OffChan = make(chan string, 100)
	newConfig.RecChan = make(chan models.Record, 100)

	if len(newConfig.AlertMap) == 0 {
		newConfig.AlertMap = make(map[string]string)
	}

	return newConfig
}
