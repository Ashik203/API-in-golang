package apps

import (
	"app/db"
	"app/web"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	db.ConnectDb()

}

func (app *Application) Run() {
	web.RunServer(&app.wg)

}

func (app *Application) Wait() {
	app.wg.Wait()
}
