package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int32 = 800, 600

type gameState int

const (
	start gameState = iota
	play
)

var state = start

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {

	window, err := sdl.CreateWindow("Testing Sprite Animation", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, winWidth, winHeight)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	//spriteSheet, err := os.Open("test.png")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer spriteSheet.Close()

	imageData, imageType, err := image.Decode(spriteSheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	spriteSheet.Seek(0, 0)

	loadedImage, err := png.Decode(spriteSheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loadedImage)

	// What is .type(spriteSheet)?

	// What do docs say about sdl_surface type?

	//tex, err := renderer.CreateTexture()

	pixels := make([]byte, winWidth*winHeight*4)

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32

	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		if state == play {
			// Do stuff
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = play
			}
		}

		clear(pixels)
		tex.Update(nil, pixels, int(winWidth)*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < .005 {
			sdl.Delay(5 - uint32(elapsedTime/1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
