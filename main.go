package main

import "myapp/ui"

func main() {
	app := ui.NewApp()

	if err := app.Init(); err != nil {
		panic(err.Error())
	}

	defer app.Destroy()

	app.Run()
}
