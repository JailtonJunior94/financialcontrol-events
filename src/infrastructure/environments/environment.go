package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	Environment      = ""
	ConnectionString = ""
	Cron             = ""
	TelegramBaseUrl  = ""
	BotKey           = ""
	ChatId           = 0
)

func New() {
	var err error

	viper.SetConfigName(fmt.Sprintf("config.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	Environment = viper.GetString("environment")
	ConnectionString = viper.GetString("connectionString")
	Cron = viper.GetString("cron")
	TelegramBaseUrl = viper.GetString("telegramBaseUrl")
	BotKey = viper.GetString("botKey")
	ChatId = viper.GetInt("chatId")
}
