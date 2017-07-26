package main

import (
	"time"

	"github.com/polarpayne/tmptris/board"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	frameRate    = 60
	timePerFrame = time.Second / frameRate
)

var (
	winName               = "hello"
	backgroundColor       = sdl.Color{R: 210, G: 120, B: 120, A: 255}
	borderColor           = sdl.Color{R: 120, G: 120, B: 120, A: 255}
	oldWinWidth     int32 = 640
	oldWinHeight    int32 = 480
	winWidth              = oldWinWidth
	winHeight             = oldWinHeight
	playerColor           = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	playerButtons   buttons
	player          playerObject
)

type playerObject struct {
	x   int32
	y   int32
	rot int
}

func handleEvents() bool {
	playerButtons.reset()

	running := true

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_w, sdl.K_UP:
				playerButtons.up = true
			case sdl.K_d, sdl.K_RIGHT:
				playerButtons.right = true
			case sdl.K_s, sdl.K_DOWN:
				playerButtons.down = true
			case sdl.K_a, sdl.K_LEFT:
				playerButtons.left = true
			case sdl.K_j:
				playerButtons.cw = true
			case sdl.K_k:
				playerButtons.ccw = true
			}
		case *sdl.WindowEvent:
			switch t.Event {
			case sdl.WINDOWEVENT_SIZE_CHANGED, sdl.WINDOWEVENT_RESIZED:
				winWidth = t.Data1
				winHeight = t.Data2
			}
		}
	}

	return running
}

func update() {
	if playerButtons.up {
		player.y -= 10
	}
	if playerButtons.right {
		player.x += 10
	}
	if playerButtons.down {
		player.y += 10
	}
	if playerButtons.left {
		player.x -= 10
	}
	if playerButtons.cw {
		player.rot++
	}
	if playerButtons.ccw {
		player.rot--
	}
}

var pieces board.Pieces

func drawTetromino(piece string, surface *sdl.Surface) {
	i, ok := pieces[piece]
	if !ok {
		panic("töttöröö")
	}

	rot := player.rot % len(i.States)
	if rot < 0 {
		rot = -rot
	}

	for x, bs := range i.States[rot] {
		for y, b := range bs {
			if b {
				surface.FillRect(&sdl.Rect{
					X: int32(x*10) + player.x, Y: int32(y*10) + player.y,
					W: 10, H: 10},
					playerColor.Uint32())
			}
		}
	}
}

func main() {
	pieces = board.ParseFile("data/tgm_pieces.data")

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow(winName,
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		int(winWidth), int(winHeight), sdl.WINDOW_SHOWN) //|sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true

	for running {
		t0 := time.Now()

		running = handleEvents()
		update()

		if winWidth != oldWinWidth || winHeight != oldWinHeight {
			oldWinWidth = winWidth
			oldWinHeight = winHeight

			surface, err = window.GetSurface()
			if err != nil {
				panic(err)
			}
		}

		// draw bg
		surface.FillRect(&sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H}, borderColor.Uint32())
		surface.FillRect(&sdl.Rect{X: 10, Y: 10, W: surface.W - 20, H: surface.H - 20}, backgroundColor.Uint32())

		drawTetromino("J", surface)
		// surface.FillRect(&player, playerColor.Uint32())

		window.UpdateSurface()

		delta := time.Now().Sub(t0)
		// fmt.Println("delta -", delta)
		time.Sleep(timePerFrame - delta)
	}
}
