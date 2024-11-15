package main

import (
	"math/rand"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
)

type Entity struct {
	Tex rl.Texture2D
	Pos rl.Vector2
}

var (
	player *Entity
	starPositions []rl.Vector2
	metorPositions []rl.Vector2
	starTexture rl.Texture2D
	meteorTexture rl.Texture2D
)

var (
	running = true
)

func init() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Go Raylib space shooter")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)


	//init player
	player = &Entity{}
	player.Tex = rl.LoadTexture(filepath.Join("resources", "images", "player.png"))
	player.Pos = rl.Vector2{X: SCREEN_WIDTH / 2, Y: SCREEN_HEIGHT / 2}

	//stars
	starTexture = rl.LoadTexture(filepath.Join("resources", "images", "star.png"))
	for i := 0; i < 20; i++ {
		starPositions = append(starPositions, randomPos())
	}

	//meteors
	meteorTexture = rl.LoadTexture(filepath.Join("resources", "images", "meteor.png"))
	metorPositions = append(metorPositions, randomPos())

}

func randomPos() rl.Vector2 {
	x := float32(rand.Intn(SCREEN_WIDTH))
	y := float32(rand.Intn(SCREEN_HEIGHT))
	return rl.Vector2{X:x, Y:y}
}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()
	if player.Pos.X < SCREEN_WIDTH {
		player.Pos.X += 5
	} else if player.Pos.X >= SCREEN_WIDTH {
		player.Pos.X = 0
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)
	
	for _, pos := range starPositions {
		rl.DrawTexture(starTexture, int32(pos.X), int32(pos.Y), rl.White)
	}

	for _, pos := range metorPositions {
		rl.DrawTexture(meteorTexture, int32(pos.X), int32(pos.Y), rl.White)
	}

	rl.DrawTexture(player.Tex, int32(player.Pos.X), int32(player.Pos.Y), rl.White)	
		
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
