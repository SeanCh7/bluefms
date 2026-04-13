package main

import (
	"time"

	"github.com/SeanCh7/bluefms/internal"
)

func main() {
	//set up server config
	serverConfig := internal.ServerConfig{
		Addr:    ":8080",
		Timeout: 15 * time.Second,
	}

	//start webserver
	webserver := internal.New(serverConfig)

	//set up webserver base route (index.html)

}
