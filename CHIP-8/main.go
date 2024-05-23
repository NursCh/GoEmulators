package main

import (
		"os"
		"github.com/veandco/go-sdl2/sdl"
)

//Initialize SDL
func init_sdl() bool { 
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
				panic(err)
				sdl.Log("Couldn't initialize SDL")
				return false
		}
		return true
}

//Final cleanup
func final_cleanup() {
		sdl.Quit()
}
func main() {
		//Successful exit status
		//
		defer os.Exit(0)
		// Initialize SDL
		//
		if !init_sdl() {
				os.Exit(1)
		}
		//Final cleanup
		//
		defer final_cleanup()
}
