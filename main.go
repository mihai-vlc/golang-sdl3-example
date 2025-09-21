package main

import (
	"fmt"
	"log"
	"time"

	"golang.design/x/hotkey"

	"github.com/jupiterrider/purego-sdl3/sdl"
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
			fmt.Println(event.Edit().Start)
			fmt.Println(event.Edit().Length)

		case sdl.EventKeyDown:
			if event.Key().Scancode == sdl.ScancodeEscape {
				running = false
			}
		}

		if isGlobalHotkeyPressed {
			fmt.Println("pressed global hotkey")
		}

		sdl.SetRenderDrawColor(renderer, 100, 150, 200, 255)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderDebugText(renderer, 100, 100, "Hello")

		var rect = &sdl.FRect{X: 0, Y: 0, W: 100, H: 100}
		sdl.RenderRect(renderer, rect)

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
