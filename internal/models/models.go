package models

import (
	"time"
)

// Conf - web gui config
type Conf struct {
	DB       string
	Host     string
	Port     string
	Theme    string
	Icon     string
	ConfPath string
	LogPath  string
	YamlPath string
	AlertMap map[string]string
	Quit     chan bool
	HashChan chan string
}

// Host - host for internal use
type Host struct {
	Name     string
	Hash     string
	Interval string
	Alerts   []string
	LastSeen time.Time
	IntSec   int
	Active   bool
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
}
