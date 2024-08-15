package main

import (
	"app/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	app.Wait()
	app.Cleanup()
}
