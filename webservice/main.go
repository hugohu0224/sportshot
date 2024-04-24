package main

import (
	"sportshot/webservice/initinal"
)

func main() {
	// initialize
	initinal.InitEventServerConn()
	Router := initinal.InitRouters()

	Router.Run("0.0.0.0:8081")
}
