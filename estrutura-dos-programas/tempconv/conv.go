package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit((c * 9 / 5) + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func KToF(k Kelvin) Fahrenheit  { celsius := KToC(k); return CToF(celsius) }
