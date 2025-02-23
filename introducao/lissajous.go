// Lissajous gera animações GIF de figuras de Lissajous aleatórias
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
)

var pallete = []color.Color{color.RGBA{144, 245, 145, 8}, color.RGBA{0, 84, 245, 8}, color.Black, color.RGBA{255, 245, 145, 8}, color.RGBA{229, 9, 229, 86}}

const (
	blueIndex   = 1 // primeira cor da paleta
	blackIndex  = 2 // próxima cor da paleta
	yellowIndex = 3
	purpleIndex = 4
)

func main() {
	// rand.NewSource(time.Now().UTC().UnixNano())
	// lissajous(os.Stdout)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := r.URL.Query().Get("cycles")
		if cycles != "" {
			cyclesConvert, err := strconv.Atoi(cycles)
			if err != nil {
				log.Fatal("Não foi possível converter")
			}
			lissajous(w, cyclesConvert)
		} else {
			lissajous(w, 5)
		}
	})
	http.ListenAndServe("localhost:8000", nil)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			max := len(pallete)
			randomIndex := rand.Intn(max) + 1
			switch randomIndex {
			case 1:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
			case 2:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			case 3:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), yellowIndex)
			case 4:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), purpleIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
