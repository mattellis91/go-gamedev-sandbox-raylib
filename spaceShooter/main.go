package main

import (
	"math/rand"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
	TARGET_FPS = 60
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
	paused bool
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

	if rl.IsKeyPressed(rl.KeyEscape) {
		paused = !paused
	}

	if !paused {
		//movement
		if rl.IsKeyDown(rl.KeyA) {
			player.Pos.X -= 5
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.Pos.X += 5
		}
		if rl.IsKeyDown(rl.KeyW) {
			player.Pos.Y -= 5
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.Pos.Y += 5
		}

		//screen wrap
		if  player.Pos.X >= SCREEN_WIDTH {
			player.Pos.X = 0
		}
		if player.Pos.X + float32(player.Tex.Width) <= 0 {
			player.Pos.X = SCREEN_WIDTH - float32(player.Tex.Width)
		}
		if player.Pos.Y < 0 {
			player.Pos.Y = SCREEN_HEIGHT
		}
		if player.Pos.Y > SCREEN_HEIGHT {
			player.Pos.Y = 0
		}
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

	if paused {
		rl.DrawText("PAUSED", SCREEN_WIDTH / 2, SCREEN_HEIGHT / 2, 20, rl.White)
	}
		
	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(starTexture)
	rl.UnloadTexture(player.Tex)
	rl.UnloadTexture(meteorTexture)
	rl.CloseWindow()
}

func main() {

	rl.SetTargetFPS(TARGET_FPS);

	for running {
		input()
		update()
		draw()
	}

	quit()
}
