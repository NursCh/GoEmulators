package main

import (
		"os"
		"github.com/veandco/go-sdl2/sdl"
)

//CHIP8 Containder object
type chip8_object struct{
		chip8_window *sdl.Window
		chip8_renderer *sdl.Renderer
}

//Configuration for CHIP-8
type chip8_config struct{
		window_width int32 // SDL window width
		window_height int32 // SDL window height
}

//Initialize SDL
func init_sdl(chip8 *chip8_object,config *chip8_config) bool { 
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
				panic(err)
				sdl.Log("Couldn't initialize SDL: %s\n", err)
				return false
		}
		window, err := create_window(config)
		if err != nil {
				sdl.Log("Couldn't create window: %s\n", err)
				return false
		}
		renderer, err := create_renderer(window)
		if err != nil {
				sdl.Log("Couldn't create renderer: %s\n", err)
				return false
		}
		chip8.chip8_window = window
		chip8.chip8_renderer = renderer
		return true
}

//Final cleanup
func final_cleanup(chip8 *chip8_object) {
		if chip8.chip8_window != nil {
		chip8.chip8_window.Destroy()
}
		sdl.Quit()
}

//Create window
func create_window(config *chip8_config) (*sdl.Window, error){
		window, err := sdl.CreateWindow("CHIP-8", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, config.window_width, config.window_height,0)
		if err != nil {
				return nil, err
		}
		return window, nil
}

//Create renderer
func create_renderer(window *sdl.Window) (*sdl.Renderer, error){
		renderer, err := sdl.CreateRenderer(window, -1, 0x00000002)  // SDL_RENDERER_ACCELERATED
		if err != nil {
				return nil, err
		}
		return renderer, nil
}

//Initialize configuration
func set_config(config *chip8_config, args []string){
		//Default settings
		config = &chip8_config{
				window_width: 64,         // Original CHIP8 width
				window_height: 32,        // Original CHIP8 height
		}
		//Custom settings (WIP)
}

//main function
func main() {
		//Successful exit status
		//
		defer os.Exit(0)
		//Create CHIP-8 objects
		//
		chip8 := chip8_object{}
		config := chip8_config{}
		//Initialize config for CHIP8
		//
		set_config(&config, os.Args[1:])
		// Initialize SDL
		//
		if !init_sdl(&chip8, &config ) {
				os.Exit(1)
		}
		//Final cleanup
		//
		defer final_cleanup(&chip8)
}
