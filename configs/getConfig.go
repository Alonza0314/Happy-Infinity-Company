package configs

import (
	"errors"
	"strconv"

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

func GetSessionCookieTimeout(identifier string) (int, error) {
	timeout, err := GetConfigs(identifier)
	if err != nil {
		return 0, err
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return 0, errors.New("server error => Atoi function error\n\t" + err.Error())
	}
	return timeoutInt, nil
}

func GetGracefulTimeout() (int, error) {
	timeout, err := GetConfigs("graceful.timeout")
	if err != nil {
		return 0, err
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		return 0, errors.New("server error => Atoi function error\n\t" + err.Error())
	}
	return timeoutInt, nil
}
