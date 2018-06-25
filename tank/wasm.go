// +build js,wasm

// from: https://codepen.io/jadch/pen/MbEabY
package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"

	"github.com/Chyroc/web"
)

func init() {
	rand.Seed(time.Now().Unix())

	web.Document.AddEventListener("keydown", func(args []js.Value) {
		if len(args) > 0 {
			keyCode := args[0].Get("keyCode").Int()
			switch keyCode {
			case 39:
				game.right_pressed = true
			case 37:
				game.left_pressed = true
			case 32:
				game.space_pressed = true
			}
		}
	})

	web.Document.AddEventListener("keyup", func(args []js.Value) {
		if len(args) > 0 {
			keyCode := args[0].Get("keyCode").Int()
			switch keyCode {
			case 39:
				game.right_pressed = false
			case 37:
				game.left_pressed = false
			case 32:
				game.space_pressed = false
			}
		}
	})

	game = newTankGame()
}

func timenow() int64 {
	return time.Now().UnixNano() / 100000
}

var game *tankGame

type block struct {
	x      float64
	y      float64
	width  int
	height int
}

type tankGame struct {
	canvas web.HTMLCanvasElement
	ctx    web.CanvasRenderingContext2D

	player_lives int
	score        int

	tank_width  int
	tank_height int
	tank_speed  float64 //actually used to increment the Y coords of all the blocks, set at 0.3

	tank block

	balls           []block
	ball_speed      float64
	since_last_fire int64

	blocks                    []block
	tank_block_collision_bool bool

	monsters      []block
	monster_speed float64 //speed at which the monsters move along the y axis.

	//This part handles the pressing of keys used to move the tank and fire
	right_pressed bool
	left_pressed  bool
	space_pressed bool
}

func newTankGame() *tankGame {
	t := tankGame{
		player_lives: 3,
		score:        0,

		tank_width:  30,
		tank_height: 40,
		tank_speed:  0.55,

		ball_speed: 1.6,

		since_last_fire: timenow(),

		tank_block_collision_bool: true,

		monster_speed: 0.6,
	}
	t.canvas = web.Document.GetElementById("canvas").(web.HTMLCanvasElement)
	t.ctx = t.canvas.GetContext("2d")
	t.tank = block{float64(t.canvas.Width()) / 2, float64(t.canvas.Height()) - 70, t.tank_width, t.tank_height}

	return &t
}

//Drawing the Tank
func (r *tankGame) drawTank() {
	tank := r.tank

	r.ctx.BeginPath()
	r.ctx.Rect(tank.x, tank.y, tank.width, tank.height)
	r.ctx.SetFillStyle("green")

	r.ctx.Rect(float64(tank.x)+float64(tank.width)/float64(2)-float64(5), tank.y-15, 10, 15)
	r.ctx.SetFillStyle("green")

	r.ctx.Fill()
	r.ctx.ClosePath()
}

//Drawing the border blocks
func (r *tankGame) drawBorder() {
	canvas := r.canvas

	r.ctx.BeginPath()
	r.ctx.Rect(0, 0, 80, canvas.Height())
	r.ctx.Rect(float64(canvas.Width()-80), 0, 80, canvas.Height())
	r.ctx.SetFillStyle("grey")
	r.ctx.Fill()
	r.ctx.ClosePath()
}

//Initializing a new ball (starting position), adding it to a list
func (r *tankGame) drawNewBall(ball_X, ball_Y float64) {
	r.ctx.BeginPath()
	r.ctx.Arc(ball_X, ball_Y, 5, 0, math.Pi*2)
	r.balls = append(r.balls, block{ball_X, ball_Y, 3, 3})
	r.since_last_fire = timenow()
}

//drawing all of the balls of the list
func (r *tankGame) drawBalls() {
	for i := 0; i < len(r.balls); i++ {
		r.ctx.BeginPath()
		r.ctx.Arc(r.balls[i].x, r.balls[i].y, 5, 0, math.Pi*2)
		r.ctx.SetFillStyle("red")
		r.ctx.Fill()
		r.ctx.ClosePath()
	}
}

//Generates coordinates with 80 < x < (canvas.width - 120) and -260 < y < -60. Returns the coordinates.
func (r *tankGame) generateCoords() []float64 {
	var x = rand.Float64()*float64(r.canvas.Width()-80) + 80
	for int(x)+120 > r.canvas.Width() {
		x = rand.Float64()*float64(r.canvas.Width()-80) + 80
	}

	y := rand.Float64()*float64(-260-60) - 60

	return []float64{x, y}
}

//Checks that the distance between TWO blocks is greater than 140 and the difference along the Y-axis is greater than 40
func (r *tankGame) distanceCheck(X1, Y1, X2, Y2 float64) bool {
	distance := math.Sqrt(math.Pow(X1-X2, 2) + math.Pow(Y1-Y2, 2))
	if distance > 140 && math.Abs(Y1-Y2) > 40 {
		return true
	} else {
		return false
	}
}

//function that returns TRUE if we need to generate new coords because the blocks are too close, FALSE otherwise
func (r *tankGame) blockDistanceChecker(X, Y float64) bool {
	if len(r.blocks) == 0 {
		return false
	}

	var check = false
	for i := 0; i < len(r.blocks); i++ {
		if r.distanceCheck(X, Y, r.blocks[i].x, r.blocks[i].y) {
			check = check || false
		} else {
			check = check || true
		}
	}

	if !check {
		return false
	} else {
		return true
	}
}

//Creates a new block
func (r *tankGame) drawNewBlock() {
	var coords = r.generateCoords()
	var X = coords[0]
	var Y = coords[1]
	for r.blockDistanceChecker(X, Y) {
		coords = r.generateCoords()
		X = coords[0]
		Y = coords[1]
	}

	var width = 40
	var height = 60
	r.blocks = append(r.blocks, block{X, Y, width, height})
}

func (r *tankGame) drawBlocks() {
	ctx := r.ctx
	for i := 0; i < len(r.blocks); i++ {
		ctx.BeginPath()
		ctx.Rect(r.blocks[i].x, r.blocks[i].y, r.blocks[i].width, r.blocks[i].height)
		ctx.SetFillStyle("green")
		ctx.Fill()
		ctx.ClosePath()
	}
}

//Mover function: moves blocks elements down (y++) along the y axis
func (r *tankGame) moverFunc() {
	for i := 0; i < len(r.blocks); i++ {
		r.blocks[i].y = r.blocks[i].y + r.tank_speed
		//Drops the block from the blocks array when they're out of view
		if r.blocks[i].y > float64(r.canvas.Width()) {
			r.blocks = append(r.blocks[:i], r.blocks[i+1:]...)
		}
	}
}

//Mover function: moves the balls the balls up (y--) along the y axis //ADDED: Moving the monsters
func (r *tankGame) moveBalls() {
	//Moving the Balls
	for i := 0; i < len(r.balls); i++ {
		r.balls[i].y = r.balls[i].y - r.ball_speed
		//Drops the ball from the balls array when they're out of view
		if r.balls[i].y < 0 {
			r.balls = append(r.balls[:i], r.balls[i+1:]...)
		}
	}

	//Moving the Monsters
	for i := 0; i < len(r.monsters); i++ {
		r.monsters[i].y = r.monsters[i].y + r.monster_speed
		//Drops the monsters from the list when they're out of view
		if r.monsters[i].y > float64(r.canvas.Width()) {
			r.monsters = append(r.monsters[:i], r.monsters[i+1:]...)
		}
	}
}

//Block Collision detection function
func (r *tankGame) tank_block_collision() {
	for i := 0; i < len(r.blocks); i++ {
		var conflict_X = false
		var conflict_Y = false
		if r.tank.x+float64(r.tank_width) > r.blocks[i].x && r.tank.x < r.blocks[i].x+40 {
			conflict_X = conflict_X || true
		}
		if r.tank.y < r.blocks[i].y+60 && r.tank.x > r.blocks[i].y {
			conflict_Y = conflict_Y || true
		}
		if conflict_X && conflict_Y {
			r.tank_block_collision_bool = false
			r.player_lives -= 1
			return
		}
	}
	r.tank_block_collision_bool = true
}

//Generates X and Y coordinates for a new monster
func (r *tankGame) create_monster() {
	var coords = r.generateCoords()
	var X = coords[0]
	var Y = coords[1]
	r.monsters = append(r.monsters, block{X, Y, 25, 29})
}

//Draws a monster
func (r *tankGame) draw_monster(X, Y float64) {
	var scale = 0.8
	var h float64 = 9 //height
	var a float64 = 5
	ctx := r.ctx

	ctx.BeginPath()
	//First trapezoid
	ctx.MoveTo(X, Y)
	ctx.LineTo(X-a*scale, Y+h*scale)
	ctx.LineTo(X+(a*4)*scale, Y+h*scale)
	ctx.LineTo(X+(a*3)*scale, Y)
	//Second trapezoid
	ctx.MoveTo(X-(a+5)*scale, Y+h*scale)
	ctx.LineTo(X-(a)*scale, Y+(h+20)*scale)
	ctx.LineTo(X+(a+15)*scale, Y+(h+20)*scale)
	ctx.LineTo(X+(a+20)*scale, Y+(h)*scale)
	ctx.SetFillStyle("purple")
	ctx.Fill()
	ctx.ClosePath()
}

//Draws all monsters in the monsters list
func (r *tankGame) draw_monsters() {
	for i := 0; i < len(r.monsters); i++ {
		var X = r.monsters[i].x
		var Y = r.monsters[i].y
		r.draw_monster(X, Y)
	}
}

//Collision detector: detects a collision along the Y-axis between two objects (Maps) with the following template: ["X", "Y", "width", "height"]
func (r *tankGame) collision_detector(first, second block) bool {
	var x1 = first.x
	var y1 = first.y
	var width1 = float64(first.width)
	var height1 = float64(first.height)
	var x2 = second.x
	var y2 = second.y
	var width2 = float64(second.width)
	var height2 = float64(second.height)

	if x2 > x1 && x2 < x1+width1 || x1 > x2 && x1 < x2+width2 {
		if y2 > y1 && y2 < y1+height1 || y1 > y2 && y1 < y2+height2 {
			return true
		}
	}
	return false
}

//detects a collion betweenn balls and monsters
func (r *tankGame) ball_monster_collision() {
	for i := 0; i < len(r.monsters); i++ {
		var monster = r.monsters[i]
		for j := 0; j < len(r.balls); j++ {
			var ball = r.balls[j]
			if r.collision_detector(monster, ball) {
				r.balls = append(r.balls[:i], r.balls[i+1:]...)
				r.monsters = append(r.monsters[:i], r.monsters[i+1:]...)
				r.score += 2
			}
		}
	}
}

//Used to display the number of lives remaining and game score
func (r *tankGame) drawInfo() {
	ctx := r.ctx
	ctx.SetFont("bold 15px Gill Sans MT")
	ctx.SetFillStyle("blue")
	ctx.FillText("Lives: "+strconv.Itoa(r.player_lives), 635, 22)
	ctx.FillText("Score: "+strconv.Itoa(r.score), 13, 22)
}

//Main function, will be used to run the game
func draw() {
	game.ctx.ClearRect(0, 0, game.canvas.Width(), game.canvas.Height())
	game.drawBorder()
	game.drawInfo()
	game.drawTank()
	game.drawBalls()
	game.drawBlocks()
	game.draw_monsters()
	game.moveBalls()
	game.tank_block_collision()
	game.ball_monster_collision()
	game.moverFunc()
	if game.space_pressed && len(game.balls) < 10 && timenow()-game.since_last_fire > 500 {
		game.drawNewBall(game.tank.x+15, game.tank.y-30)
	}
	if game.right_pressed && game.tank.x+float64(game.tank_width) < float64(game.canvas.Width()) {
		game.tank.x = game.tank.x + 1
	}
	if game.left_pressed && game.tank.x > 0 {
		game.tank.x = game.tank.x - 1
	}

	if len(game.blocks) < 3 {
		game.drawNewBlock()
	}
	if len(game.monsters) < 1 {
		game.create_monster()
	}
	if !game.tank_block_collision_bool && game.player_lives < 0 {
		web.Window.Alert("you lost.")
		web.Document.Location().Reload()
		return
	}

	web.Window.RequestAnimationFrame(js.NewCallback(func(args []js.Value) {
		draw()
	}))
}

func main() {
	fmt.Printf("tank %v\n", game.tank)

	draw()

	select {}
}
