package main

import (
	"fmt"
	"log"
	"time"

	"golang.design/x/hotkey"
)

func registerGlobalHotkey() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyD)
	err := hk.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}

	log.Printf("hotkey: %v is registered\n", hk)

	go func() {
		for {
			select {
			case <-hk.Keydown():
				t := time.Now()
				fmt.Printf("keydown %s\n", t.Format(time.RFC850))
			case <-hk.Keyup():
				t := time.Now()
				fmt.Printf("keyup %s\n", t.Format(time.RFC850))
			}
		}
	}()

	// hk.Unregister()
	// log.Printf("hotkey: %v is unregistered\n", hk)
}
