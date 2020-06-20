package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int32 = 800, 600

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

	spriteSheet, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spriteSheet.Close()

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
}
