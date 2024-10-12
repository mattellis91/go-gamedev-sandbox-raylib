package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
)

var (
	running = true
)

func init() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Go Raylib space shooter")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)
	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		draw()
	}

	quit()
}