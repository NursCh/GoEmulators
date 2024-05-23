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
func final_cleanup() int {
		sdl.Quit()
		return 0
}
func main() {
		// Initialize SDL
		//
		if !init_sdl() {
				os.Exit(1)
		}
		//Final cleanup
		//
		os.Exit(final_cleanup())
}
