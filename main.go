package main

import (
	"github.com/provsalt/beast-backend/config"
	"github.com/provsalt/beast-backend/webserver"
)

func main() {
	webserver.New(config.LoadConfig())
}
