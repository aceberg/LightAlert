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
}

// Host - host to monitor
type Host struct {
	Name     string   `yaml:"name"`
	Hash     string   `yaml:"hash"`
	Interval string   `yaml:"interval"`
	Alerts   []string `yaml:"alerts"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
}
