package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type seed []bool

func main() {
	stitchLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	seedRand := int64(0)
	if len(os.Args) > 4 {
		giveSeed, err := strconv.Atoi(os.Args[4])
		if err != nil {
			panic(err)
		}
		seedRand = int64(giveSeed)
	}
	if seedRand == 0 {
		seedRand = time.Now().UnixNano()

	}

	fmt.Printf("using seed: %v\n", seedRand)
	// seedH := seed{true, false, true, false, true, true, true, false, true, false, false, true, false, true, false, true, false, true, false, true, false, true, false, true}
	// seedV := seed{true, false, true, false, true, false, true, true, false, true, false, true, false, false, false, true, false, true, false, true, false, true, false, true}
	// seedH := seed{true, true, false, false}
	// seedV := seed{true, true, false, false}

	seedH := randSeed(seedRand, 100, .5)
	seedV := randSeed(seedRand, 100, .5)
	img := makeImage(seedH, seedV, stitchLength, width, height)

	out, err := os.Create("stitch.png")
	if err != nil {
		panic(err)
	}
	png.Encode(out, img)
	out.Close()
}

// return a patern seed given odds of true (0-1)
func randSeed(randSeed int64, size int, oddsTrue float32) seed {
	rand.Seed(randSeed)

	seed := seed{}
	for i := 0; i < size; i++ {
		if rand.Float32() < oddsTrue {
			seed = append(seed, true)
		} else {
			seed = append(seed, false)
		}
	}
	return seed
}

// Create the image from given seeds and size.
func makeImage(seedHor seed, seedVer seed, stitchLen int, width int, height int) *image.Alpha {

	seedCols := normSeed(seedVer, width)
	seedRows := normSeed(seedHor, height)

	img := image.NewAlpha(image.Rect(0, 0, stitchLen*width, stitchLen*height))
	for row, rowSeed := range seedRows {
		addRow(img, stitchLen, row, rowSeed)
	}
	for col, colSeed := range seedCols {
		addCol(img, stitchLen, col, colSeed)
	}
	return img
}

// Add a row to the image based on the given seed.
func addRow(img *image.Alpha, stitchLen int, row int, rowSeed bool) {
	if rowSeed {
		for i := 0; i < img.Rect.Max.X; i++ {
			if i/stitchLen%2 == 0 {
				img.SetAlpha(i, row*stitchLen, color.Alpha{255})
				img.SetAlpha(i+1, row*stitchLen, color.Alpha{255})
			}
		}
	} else {
		for i := 0; i < img.Rect.Max.X; i++ {
			if i/stitchLen%2 == 1 {
				img.SetAlpha(i, row*stitchLen, color.Alpha{255})
				img.SetAlpha(i+1, row*stitchLen, color.Alpha{255})
			}
		}
	}
}

// Add a col to the image based on the given seed.
func addCol(img *image.Alpha, stitchLen int, col int, colSeed bool) {
	if colSeed {
		for i := 0; i < img.Rect.Max.Y; i++ {
			if i/stitchLen%2 == 0 {
				img.SetAlpha(col*stitchLen, i, color.Alpha{255})
				img.SetAlpha(col*stitchLen, i+1, color.Alpha{255})
			}
		}
	} else {
		for i := 0; i < img.Rect.Max.Y; i++ {
			if i/stitchLen%2 == 1 {
				img.SetAlpha(col*stitchLen, i, color.Alpha{255})
				img.SetAlpha(col*stitchLen, i+1, color.Alpha{255})
			}
		}
	}
}

// Make sure the seed is a certain length.
// If too long it will be cut.
// if to short it will be repeated.
func normSeed(s seed, length int) seed {
	ns := make(seed, length)
	for i := 0; i < length; i++ {
		ns[i] = s[i%len(s)]
	}
	return ns
}
