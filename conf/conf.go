package conf

import (
	"CodeArena/consts"
	"github.com/spf13/viper"
	"log"
)

var V *viper.Viper

func init() {
	V = viper.New()
	V.SetConfigName("config")
	V.AddConfigPath("conf")
	V.SetConfigType("toml")

	// default config
	V.SetDefault("server.port", 8080)
	V.SetDefault("server.logPath", consts.DefaultLogPath)

	err := V.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	SetLogger()
}
