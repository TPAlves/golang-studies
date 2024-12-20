package test

import (
	"testing"

	"capitulo02/pointer"
)

func TestAddress(t *testing.T) {
	valorXPTO := 10
	resultado := pointer.Address(&valorXPTO)
	esperado := &valorXPTO

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}
