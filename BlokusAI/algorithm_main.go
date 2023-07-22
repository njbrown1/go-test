package main

import (
	"fmt"
)

func main() {
	fmt.Println("running")
	pixelgl.Run(run)
}

type GameBoard struct {
	width  int
	height int
}

func NewGameBoard(width int, height int) *GameBoard {
	board := &GameBoard{
		width:  width,
		height: height,
	}
	return board
}

func drawGameBoard(win *pixelgl.Window, board *GameBoard) {
	// Calculate the size of each cell based on the window dimensions and the board size
	cellWidth := win.Bounds().W() / float64(board.width)
	cellHeight := win.Bounds().H() / float64(board.height)

	// Loop through each cell in the board and draw rectangles to represent the squares
	for x := 0; x < board.width; x++ {
		for y := 0; y < board.height; y++ {
			// Calculate the position of the current cell's top-left corner
			posX := float64(x) * cellWidth
			posY := float64(y) * cellHeight

			cellColor := colornames.Black
			// Draw a rectangle to represent the current cell
			cellRect := pixel.R(posX, posY, posX+cellWidth, posY+cellHeight)
			rect := pixel.NewSprite(nil, cellRect)
			rect.Draw(win, pixel.IM)
		}
	}
}

func run() {
	window := pixelgl.WindowConfig{
		Title:  "Blokus AI",
		Bounds: pixel.R(0, 0, 800, 800),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(window)
	if err != nil {
		panic(err)
	}

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
