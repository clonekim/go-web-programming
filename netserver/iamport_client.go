package netserver

import (
	"bonjour/iamport"
	"fmt"
)

var Client *iamport.Client

func InitClient() {
	fmt.Println("--- Initialize Client ---")

	Client = &iamport.Client{
		Uri:       Conf.Iamport.Uri,
		Mid:       Conf.Iamport.Mid,
		Secret:    Conf.Iamport.Secret,
		ImpKey:    Conf.Iamport.ImpKey,
		ImpSecret: Conf.Iamport.ImpSecret,
		Http:      iamport.NewHttpClient(Conf.Debug),
	}
}
