package configs

import (
	"errors"

	"github.com/spf13/viper"
)

func GetConfigs(identifier string) (string, error) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("configs/config.conf")
	if err := viper.ReadInConfig(); err != nil {
		return "", errors.New("server error => getconfig function error\n\t" + err.Error())
	}
	return viper.GetString(identifier), nil
}

func GetActionURL(identifier, path string) (string, error) {
	addr, err := GetConfigs(identifier)
	if err != nil {
		return "", err
	}
	flag, err := GetConfigs("security.https")
	if err != nil {
		return "", err
	}
	if flag == "true" {
		return "https://" + addr + path, nil
	}
	return "http://" + addr + path, nil
}
