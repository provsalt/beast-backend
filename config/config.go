package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config is the configuration structure
type Config struct {
	Notifier struct {
		Email struct {
			Enabled       bool   `yaml:"enabled"`
			MailgunDomain string `yaml:"mailgun_domain"`
			MailgunAPIKey string `yaml:"mailgun_api_key"`
			Domain        string `yaml:"domain"`
		}
	}
	Webserver struct {
		Host string
	}
}

// LoadConfig loads the configuration file required for the program to function
func LoadConfig() Config {
	cfg := Config{}
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		data, err := yaml.Marshal(DefaultConfig())

		if err := ioutil.WriteFile("config.yml", data, 0644); err != nil {
			fmt.Println("")
			return DefaultConfig()
		}
		if err != nil {
			return DefaultConfig()
		}
		return DefaultConfig()
	}
	f, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// DefaultConfig returns the default configuration data
func DefaultConfig() Config {
	cfg := Config{}
	cfg.Notifier.Email.Enabled = true
	cfg.Notifier.Email.MailgunAPIKey = "APIKEY"
	cfg.Notifier.Email.MailgunDomain = "email.example.com"
	cfg.Notifier.Email.Domain = "example.com"
	return cfg
}
