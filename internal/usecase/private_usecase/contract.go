package private_usecase

import tele "gopkg.in/telebot.v3"

type PrivateUsecase interface {
	Sessions(c tele.Context)
}
