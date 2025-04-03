package config

import (
	"gopkg.in/ini.v1"
)

func LoadConfigValue(section, key string) (string, error) {
	cfg, err := ini.LooseLoad("config.ini")
	if err != nil {
		return "", err
	}

	return cfg.Section(section).Key(key).String(), nil
}

func SaveConfigValue(section, key, value string) error {
	cfg, err := ini.LooseLoad("config.ini")
	if err != nil {
		return err
	}

	cfg.Section(section).Key(key).SetValue(value)
	return cfg.SaveTo("config.ini")
}
