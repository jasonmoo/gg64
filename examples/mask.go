package main

import (
	"log"

	gg "github.com/jasonmoo/gg64"
)

func main() {
	im, err := gg.LoadImage("examples/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}
