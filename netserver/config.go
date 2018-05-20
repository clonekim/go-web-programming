package netserver

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Database struct {
		Driver     string
		Connection string
		Debug      bool
	}

	Sms struct {
		Apikey    string
		Apisecret string
		Senderkey string
	}

	Google struct {
		Shortener string
	}

	Postal struct {
		Url       string
		AccessKey string
	}

	Iamport struct {
		Uri       string
		Mid       string
		Secret    string
		ImpKey    string
		ImpSecret string
	}

	Host  string
	Port  string
	Debug bool
}

var Conf Config

func LoadConfig() {
	env := os.Getenv("GOENV")

	var confile string

	if env == "" {
		confile = "config.dev.yml"
	} else if env == "prod" {
		confile = "config.yml"
	}

	box, err := rice.FindBox("yml")

	if err != nil {
		panic(err)
	}

	file, err := box.Open(confile)

	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	defer file.Close()
	viper.MergeConfig(file)

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	fmt.Printf("--- Load config from %s ---\n", confile)

}
