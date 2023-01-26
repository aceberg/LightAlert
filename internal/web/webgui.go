package web

import (
	"log"
	"net/http"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/conf"
	"github.com/aceberg/LightAlert/internal/models"
)

// Gui - start web server
func Gui(config models.Conf) {
	AppConfig = mergeConfig(config)

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	// Temp
	var oneHost models.Host
	oneHost.Name = "Test Host"
	oneHost.Hash = "fSzPfpZLUtXZgcZw5eEc"
	AllHosts = append(AllHosts, oneHost)

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/check/", checkHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}

func mergeConfig(config models.Conf) models.Conf {

	newConfig := conf.Get(config.ConfPath)
	newConfig.ConfPath = config.ConfPath

	if config.YamlPath != "" {
		newConfig.YamlPath = config.YamlPath
	}
	if config.LogPath != "" {
		newConfig.LogPath = config.LogPath
	}

	newConfig.Icon = Icon

	return newConfig
}
