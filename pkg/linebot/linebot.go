package linebot

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LineBot struct {
		ChannelSecret string `yaml:"LINE_CHANNEL_SECRET"`
		ChannelToken  string `yaml:"LINE_CHANNEL_ACCESS_TOKEN"`
	} `yaml:"line_bot"`
}

var lineBot *linebot.Client

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func InitLineBot() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	bot, err := linebot.New(config.LineBot.ChannelSecret, config.LineBot.ChannelToken)
	if err != nil {
		log.Fatal("Error initializing LINE Bot:", err)
	}

	lineBot = bot
	fmt.Println("LINE Bot initialized successfully")
}

// GetLineBot returns the LINE Bot client instance
func GetLineBot() *linebot.Client {
	return lineBot
}
