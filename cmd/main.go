package main

import (
	"collector-telegram-bot/config"
	"collector-telegram-bot/internal/server"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		configPath string
		token      string
	)
	flag.StringVar(&configPath, "c", "./config/config.toml", "path to config file")
	flag.StringVar(&token, "t", "", "token for bot")
	flag.Parse()

	serverConfig := config.CreateConfigForServer()
	_, err := toml.DecodeFile(configPath, &serverConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	contextLogger := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	appServer := server.CreateServer(serverConfig, contextLogger, token)

	appServer.Start()
}
