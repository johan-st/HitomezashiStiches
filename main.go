package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type point struct {
	x int
	y int
}

type stitch struct {
	start point
	end   point
}
type pattern []stitch

func main() {
	pat := pattern{}
	pat = append(pat, newStitch(0, 0, 0, 1))
	pat = append(pat, newStitch(2, 0, 2, 1))
	pat = append(pat, newStitch(4, 0, 4, 1))
	pat = append(pat, newStitch(6, 0, 6, 1))
	pat = append(pat, newStitch(8, 0, 8, 1))
	pat = append(pat, newStitch(0, 4, 0, 5))
	pat = append(pat, newStitch(2, 4, 2, 5))
	pat = append(pat, newStitch(4, 4, 4, 5))
	pat = append(pat, newStitch(6, 4, 6, 5))
	pat = append(pat, newStitch(8, 4, 8, 5))

	img := image.NewAlpha(image.Rect(0, 0, 30, 30))

	addPattern(img, pat)

	out, err := os.Create("stitch.png")
	if err != nil {
		panic(err)
	}
	png.Encode(out, img)
	out.Close()
}

func addPattern(i *image.Alpha, p pattern) {
	for _, s := range p {
		addHorizontalStitch(i, s)
	}
}
func addHorizontalStitch(img *image.Alpha, s stitch) {
	const scale = 3
	for i := 0; i < scale; i++ {
		img.SetAlpha(s.start.x*scale+i, s.start.y, color.Alpha{A: 255})
	}
}
func addVerticalStitch(img *image.Alpha, s stitch) {
	const scale = 3
	for i := 0; i < scale; i++ {
		img.SetAlpha(s.start.x, s.start.y*scale+i, color.Alpha{A: 255})
	}
}

func newStitch(xStart int, yStart int, xEnd int, yEnd int) stitch {
	return stitch{start: point{x: xStart, y: yStart}, end: point{x: xEnd, y: yEnd}}

}
