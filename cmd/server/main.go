package main

import (
	"app/apps"
)

func main() {
	app := apps.NewApplication()
	// app.Init()
	app.Run()
	app.Wait()
}
