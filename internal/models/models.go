package models

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
	Quit     chan bool
}

// Host - host from yaml file
type Host struct {
	Name     string `yaml:"name"`
	Hash     string `yaml:"hash"`
	Interval string `yaml:"interval"`
	LastSeen string
	Alerts   []string `yaml:"alerts"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
}
