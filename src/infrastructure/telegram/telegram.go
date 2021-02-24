package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/dtos"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/environments"
)

type ITelegram interface {
	SendMessage(message string) error
}

type Telegram struct{}

func NewTelegramService() ITelegram {
	return &Telegram{}
}

func (t *Telegram) SendMessage(message string) error {
	request := dtos.NewSendMessage(environments.ChatId, message)
	reqBytes, err := json.Marshal(&request)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s/sendMessage", environments.TelegramBaseUrl, environments.BotKey)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
