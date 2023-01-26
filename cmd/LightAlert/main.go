package main

import (
	"flag"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/web"
)

const confPath = "/data/LightAlert/config.yaml"
const logPath = ""
const yamlPath = ""

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	logPtr := flag.String("l", logPath, "Path to log file")
	yamlPtr := flag.String("h", yamlPath, "Path to hosts yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr
	conf.YamlPath = *yamlPtr
	conf.LogPath = *logPtr

	check.Path(conf.ConfPath)
	check.Path(conf.YamlPath)
	check.Path(conf.LogPath)

	web.Gui(conf)
}
