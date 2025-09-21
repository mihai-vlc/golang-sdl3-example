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

/*
func main() {
	registerGlobalHotkey()
	isGlobalHotkeyPressed := false

	app := ui.NewApp()

	if err := app.Init(); err != nil {
		panic(err.Error())
	}

	defer app.Destroy()

	if !ttf.Init() {
		panic(sdl.GetError())
	}

	font := ttf.OpenFont("resources/Roboto-Light.ttf", 80)

	if font == nil {
		panic(sdl.GetError())
	}
	defer ttf.CloseFont(font)

	textColor := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	helloSurface := ttf.RenderTextBlended(font, "Hello There", 0, textColor)
	if helloSurface == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroySurface(helloSurface)

	helloTexture := sdl.CreateTextureFromSurface(renderer, helloSurface)
	if helloTexture == nil {
		panic(sdl.GetError())
	}

	var running = true

	sdl.StartTextInput(window)

	for running {
		var event sdl.Event

		sdl.WaitEventTimeout(&event, 1000)

		switch event.Type() {
		case sdl.EventQuit:
			running = false

		case sdl.EventTextInput:
			var text = event.Text()
			fmt.Println(text.Text())

		case sdl.EventTextEditing:
			var edit = event.Edit()
			fmt.Println(edit.Text())

		case sdl.EventKeyDown:
			if event.Key().Scancode == sdl.ScancodeEscape {
				running = false
			}
		}

		if isGlobalHotkeyPressed {
			fmt.Println("pressed global hotkey")
		}

		var helloTextLocation = &sdl.FRect{}
		var w int32
		var h int32
		sdl.GetRenderOutputSize(renderer, &w, &h)
		sdl.GetTextureSize(helloTexture, &helloTextLocation.W, &helloTextLocation.H)
		helloTextLocation.X = (float32(w) - helloTextLocation.W) / 2
		helloTextLocation.Y = (float32(h) - helloTextLocation.H) / 2

		sdl.SetRenderDrawColor(renderer, 100, 150, 200, 255)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderTexture(renderer, helloTexture, nil, helloTextLocation)

		sdl.RenderPresent(renderer)
	}
}
*/
