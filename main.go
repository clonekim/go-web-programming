package main

import (
	"bonjour/netserver"
	"github.com/GeertJohan/go.rice"
)

type Asset struct {
	Template *rice.HTTPBox
}

func main() {
	asset := netserver.NewAsset("templates")

	netserver.LoadConfig()
	netserver.OpenDatabase()
	netserver.InitClient()
	netserver.EchoStart(asset)
}
