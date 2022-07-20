package main

import (
	"github.com/mochganjarn/go-template-project/app"
	"github.com/mochganjarn/go-template-project/config"
)

func main() {
	app.Run(config.InitConfig())
}
