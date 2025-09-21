package main

import (
	"myapp/ui"
)

func main() {
	app := ui.NewApp()

	if err := app.Init(); err != nil {
		panic(err.Error())
	}

	app.Run() // Main loop

	app.Destroy()
}
