package main

import (
	"fmt"
	"log"
	"time"

	"golang.design/x/hotkey"

	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
)

func main() {
	globalHotKey := registerGlobalHotkey()
	isGlobalHotkeyPressed := false

	go func() {
		for {
			select {
			case <-globalHotKey.Keydown():
				t := time.Now()
				fmt.Printf("keydown %s\n", t.Format(time.RFC850))
			case <-globalHotKey.Keyup():
				t := time.Now()
				fmt.Printf("keyup %s\n", t.Format(time.RFC850))
			}
		}
	}()

	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("Hello, World!", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyRenderer(renderer)
	defer sdl.DestroyWindow(window)

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

func registerGlobalHotkey() *hotkey.Hotkey {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyD)
	err := hk.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return nil
	}

	log.Printf("hotkey: %v is registered\n", hk)
	return hk

	// hk.Unregister()
	// log.Printf("hotkey: %v is unregistered\n", hk)
}
