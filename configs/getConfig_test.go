package configs_test

import (
	"errors"
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
)

func GetConfigs(identifier string) (string, error) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		return "", errors.New("server error => getconfig function error ")
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

func TestGetConfig(t *testing.T) {
	addr, err := GetConfigs("HICserver.addr")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, addr, "hic.alonza.xyz")
}

func TestGetActionURL(t *testing.T) {
	actionURLSignup, err := GetActionURL("HICserver.addr", "/signup")
	if err != nil {
		t.Fatal(err)
	}
	flag, err := GetConfigs("security.https")
	if err != nil {
		t.Fatal(err)
	}
	protocol := ""
	if flag == "true" {
		protocol = "https://"
	} else {
		protocol = "http://"
	}
	assert.Equal(t, actionURLSignup, protocol+"hic.alonza.xyz/signup")
}
