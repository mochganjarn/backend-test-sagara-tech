package app

import (
	"github.com/mochganjarn/go-template-project/config"
	"github.com/mochganjarn/go-template-project/router"
	"github.com/mochganjarn/go-template-project/service"
)

func Run(appConfig *config.Config) {
	router.Init(service.InstantiateDependencies(appConfig))
}
