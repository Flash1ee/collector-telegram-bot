package main

import (
	"collector-telegram-bot/config"
	"collector-telegram-bot/internal/server"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var configPath = "../config/config.toml" //TODO: Replace to Getenv

func main() {
	serverConfig := config.CreateConfigForServer()
	_, err := toml.DecodeFile(configPath, &serverConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	contextLogger := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	appServer := server.CreateServer(serverConfig, contextLogger)

	appServer.Start()
}
