package vancover

// AcretoHectare (acre float32) float32 in Hectare value
func AcretoHectare(acre float32) float32 {
	return acre / 2.471
}

// AcretoSquareinch (acre float32) float32 in Square inch value
func AcretoSquareinch(acre float32) float32 {
	return acre * 6.273e+6
}

// AcretoSquarefoot (acre float32) float32 in Square Foot value
func AcretoSquarefoot(acre float32) float32 {
	return acre * 43560
}

// AcretoSquareyard (acre float32) float32 in Square yard value
func AcretoSquareyard(acre float32) float32 {
	return acre * 4840
}

// AcretoSquaremile (acre float32) float32 in Square mile
func AcretoSquaremile(acre float32) float32 {
	return acre / 640
}

// AcretoSquaremeter (acre float32) float32 in Square meter
func AcretoSquaremeter(acre float32) float32 {
	return acre / 4046.856
}

// AcretoSquarekilometer (acre float32) float32 in Square mile
func AcretoSquarekilometer(acre float32) float32 {
	return acre / 247.105
}
