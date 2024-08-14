package apps

import (
	"app/config"
	"app/db"
	"app/web"
	"app/web/utils"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	config.LoadConfig()
	config.GetConfig()
	db.InitDB()
	utils.InitValidator()
}
func (app *Application) Run() {
	web.RunServer(&app.wg)

}

func (app *Application) Wait() {
	app.wg.Wait()
}

func (app *Application) Cleanup() {
	db.CloseDB()
}
