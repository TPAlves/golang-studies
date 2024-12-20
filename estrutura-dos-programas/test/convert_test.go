package test

import (
	"testing"

	ftoc "capitulo02/ftoc"
)

func TestConvertFahrenheitToCelsius(t *testing.T) {
	var fahrenheit ftoc.Fahrenheit = 32.0
	resultado := ftoc.FahrenheitToCelsius(fahrenheit)
	var esperado ftoc.Celsius = 0.0

	if resultado != esperado {
		t.Errorf("Resultado %g°C, esperado %g°C", resultado, esperado)
	}
}
