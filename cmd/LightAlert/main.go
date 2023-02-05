package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
	"github.com/aceberg/LightAlert/internal/web"
)

const confPath = "/data/LightAlert/config.yaml"
const dbPath = ""
const yamlPath = ""

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	dbPtr := flag.String("d", dbPath, "Path to DB file")
	yamlPtr := flag.String("h", yamlPath, "Path to hosts yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr
	conf.YamlPath = *yamlPtr
	conf.DB = *dbPtr

	check.Path(conf.ConfPath)
	check.Path(conf.YamlPath)

	web.Gui(conf)
}
