package models

import (
	"time"
)

// Conf - web gui config
type Conf struct {
	DB       string
	Host     string
	Port     string
	Show     string
	Theme    string
	Icon     string
	ConfPath string
	YamlPath string
	AlertMap map[string]string
	Quit     chan bool
	HashChan chan string
	OffChan  chan string
	RecChan  chan Record
}

// Host - host for internal use
type Host struct {
	Name     string
	Hash     string
	Interval string
	Alerts   []string
	LastSeen time.Time
	IntSec   int64
	Active   bool
}

// Record - write check to DB
type Record struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Name  string `db:"NAME"`
	Hash  string `db:"HASH"`
	IP    string `db:"IP"`
	Agent string `db:"AGENT"`
	State string `db:"STATE"`
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Hosts   []Host
	Records []Record
	Len     int
}
