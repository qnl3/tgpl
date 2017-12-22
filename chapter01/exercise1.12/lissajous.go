package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var (
	palette = []color.Color{color.White, color.Black}
)

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	count.increment()

	if err := r.ParseForm(); err != nil {
		lissajous(w, 0.1)
		return
	}

	qCycles := strings.Join(r.Form["cycles"], "")
	if qCycles == "" {
		lissajous(w, 0.1)
		return
	}

	c, err := strconv.ParseFloat(qCycles, 64)

	if err != nil {
		log.Fatal(err)
	}

	lissajous(w, c)
}

func lissajous(out io.Writer, c float64) {
	var cycles = c // number of complete x oscillator revolutions
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 6     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoreing encoding errors

}
