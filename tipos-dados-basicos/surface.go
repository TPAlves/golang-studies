// Surface calcula uma renderização SVG de uma função de superfice 30

// @TODO: Criar o arquivo .SVG e permitir a passagem dinamica dos valores

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	width, height = 600, 300
	cells         = 100
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cons30 = math.Sin(angle), math.Cos(angle)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("xyrange não foi passado, argumento obrigatório.")

	}
	xyrange, err := strconv.ParseFloat(strings.TrimSpace(os.Args[1]), 64)
	if err != nil {
		log.Fatal("Argumento inválido, são aceitos somente números.")
	}

	if xyrange == 0 {
		log.Fatal("Valor inválido, não é aceito `0` como argumento.")
	}

	xyscale := width / 2 / xyrange

	fmt.Printf(`<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill: white; strokewidth: 0,7' width='%d' height='%d'> `, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(xyrange, xyscale, i+1, j)
			bx, by := corner(xyrange, xyscale, i, j)
			cx, cy := corner(xyrange, xyscale, i, j+1)
			dx, dy := corner(xyrange, xyscale, i+1, j+1)
			fmt.Printf("<polygon points='%g, %g, %g, %g %g, %g %g, %g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(xyrange, xyscale float64, i, j int) (float64, float64) {
	// Encontra o ponto (x, y) no canto da célula (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// calcula a altura z da superficie
	z := f(x, y)

	// faz uma projeção isométrica de x,y,z sobre sx, sy do canvas SVG 2D
	sx := width/2 + (x-y)*cons30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distancia de (0,0)
	return math.Sin(r) / r
}
