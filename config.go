package go_pi_monitor

import (
	"github.com/BurntSushi/toml"
	"log"
)

type AppConfig struct {
	Http struct {
		Port            string `toml:"port"`
		Favicon         string `toml:"favicon"`
		StaticPath      string `toml:"static_path"`
		HtmlPathPattern string `toml:"html_path_pattern"`
	} `toml:"http"`
}

func initConfig() *AppConfig {
	var config *AppConfig
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Fatalln(err)
	}
	return config
}
