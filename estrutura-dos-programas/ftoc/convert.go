package ftoc

type Fahrenheit float64
type Celsius float64

// Converte Fahrenheit para Celsius
func FahrenheitToCelsius(valueFahrenheit Fahrenheit) (valueCelsius Celsius) {
	valueCelsius = Celsius(valueFahrenheit - 32) * 5 / 9
	return
}
