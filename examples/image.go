package main

import "github.com/nsf/gothic"
import "image/png"
import "image"
import "os"

func loadPNG(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}

func main() {
	img := loadPNG("background.png")
	ir, err := gothic.NewInterpreter()
	if err != nil {
		panic(err)
	}

	ir.UploadImage("bg", img)
	ir.Eval(`
canvas .c -width 640 -height 480 -highlightthickness 0
.c create image 0 0 -anchor nw -image bg -tags mybg
.c lower mybg
pack .c -expand true
	`)
	ir.MainLoop()
}
