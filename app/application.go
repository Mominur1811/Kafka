package app

import (
	"kafka-by-momin/config"
	"kafka-by-momin/kafka"
	"kafka-by-momin/web"
	"kafka-by-momin/web/utils"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	// initialize application dependencies
	config.LoadConfig()
	utils.InitValidator()
	kafka.InitKafkaProducer()
}

func (app *Application) Run() {
	// run different goroutines and services in non-blocking way
	web.StartServer(&app.wg)
}

func (app *Application) Wait() {
	// wait for all the top level goroutines
	app.wg.Wait()
}

func (app *Application) Cleanup() {
	// close connections or any other memory clean up operation
	kafka.CloseKafkaProducer()

}
