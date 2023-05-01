package server

import (
	"collector-telegram-bot/config"
	"collector-telegram-bot/internal/delivery"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"time"
)

type Server struct {
	config *config.ServerConfig
	logger *logrus.Entry
	token  string
}

func CreateServer(config *config.ServerConfig, logger *logrus.Entry, token string) *Server {
	return &Server{config: config, logger: logger, token: token}
}

func (s *Server) Start() {
	pref := telebot.Settings{
		Token:  s.token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		s.logger.Fatalf("Server error: %s", fmt.Sprintf("%v", err))
	}

	privateHandler := delivery.MakePrivateTgHandler(s.logger)
	groupHandler := delivery.MakeGroupTgHandler(s.logger)

	b.Handle("/start", privateHandler.Start)
	b.Handle("/info", privateHandler.Info)
	b.Handle("@collector_money_bot start", groupHandler.Great)

	s.logger.Info("Server is working")

	b.Start()
}
