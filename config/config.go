package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

func GetConfig() *Config {
	return config
}

type Config struct {
	ChatGpt ChatGptConfig `json:"chatgpt,omitempty" mapstructure:"chatgpt" yaml:"chatgpt"`
	Wecom   WecomConfig   `json:"wecom,omitempty" mapstructure:"wecom" yaml:"wecom"`
}

type ChatGptConfig struct {
	Token string `json:"token,omitempty"  mapstructure:"token,omitempty"  yaml:"token,omitempty"`
}

type WecomConfig struct {
	RobotKey      string `json:"robot_key,omitempty" mapstructure:"robot_key,omitempty" yaml:"robot_key"`
	ReplyTemplate string `json:"reply_template,omitempty" mapstructure:"reply_template,omitempty" yaml:"reply_template"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

func GetOpenAiApiKey() *string {
	apiKey := getEnv("api_key")
	if apiKey != nil {
		return apiKey
	}

	if config == nil {
		return nil
	}

	if apiKey == nil {
		apiKey = &config.ChatGpt.Token
	}
	return apiKey
}

func GetWecomRobotKey() *string {
	robotKey := getEnv("robot_key")
	if robotKey != nil {
		return robotKey
	}

	if config == nil {
		return nil
	}

	if robotKey == nil {
		robotKey = &config.Wecom.RobotKey
	}
	return robotKey
}

func getEnv(key string) *string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = os.Getenv(strings.ToUpper(key))
	}

	if len(value) > 0 {
		return &value
	}

	if config == nil {
		return nil
	}

	if len(value) > 0 {
		return &value
	}

	return nil
}
