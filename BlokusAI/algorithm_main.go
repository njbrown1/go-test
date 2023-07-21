package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Blokus AI",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(win)

	for !win.Closed() {
		// Update game logic here
		// ...

		// Clear the window with a background color
		win.Clear(colornames.White)

		// Draw the game board and pieces here
		// ...

		// Update the window
		win.Update()
	}
}
