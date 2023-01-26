package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("DB", "/data/LightAlert/sqlite.db")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8846")
	viper.SetDefault("THEME", "cerulean")
	viper.SetDefault("YAMLPATH", "/data/LightAlert/hosts.yaml")
	viper.SetDefault("LOGPATH", "/data/LightAlert/LightAlert.log")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DB, _ = viper.Get("DB").(string)
	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.LogPath, _ = viper.Get("LOGPATH").(string)
	config.YamlPath, _ = viper.Get("YAMLPATH").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("db", config.DB)
	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("logpath", config.LogPath)
	viper.Set("yamlpath", config.YamlPath)

	err := viper.WriteConfig()
	check.IfError(err)
}