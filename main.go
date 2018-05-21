package main

import "bonjour/netserver"

func main() {

	netserver.LoadConfig()
	db := netserver.OpenDatabase()
	defer db.Close()
	netserver.SetupLogger()
	netserver.InitClient()
	netserver.EchoStart()
}
