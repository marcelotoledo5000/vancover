// The length provide you a source of unit
// convetion of lenght

package converter

const (
	kmmi      = 0.62137
	yardmeter = 1.0936
)

var (
	centimeter  float32
	feet        float32
	inches      float32
	meter       float32
	miles       float32
	kilometers  float32
	millimeters float32
	yard        float32
)

// MilestoKilometers returns Kilometers in x Miles
func MilestoKilometers(kilometers float32) float32 {
	miles = kilometers * kmmi
	return miles
}

// KilometerstoMiles return Miles in x Kilometers
func KilometerstoMiles(miles float32) float32 {
	kilometers = miles / kmmi
	return kilometers
}

// FeetoMeters converts feet to meters
func FeetoMeters(feet float32) float32 {
	feet *= 0.3048
	return feet
}

// FeetoYard (feet float32) returns flaot32 in yard value
func FeetoYard(feet float32) float32 {
	yard = feet * 0.33333
	return yard
}

// MeterstoFeet converts Meters to feet
func MeterstoFeet(meter float32) float32 {
	feet = meter * 3.2808
	return feet
}

// InchestoCentimeter (inches float32) returns float32 in Centimeter
func InchestoCentimeter(inches float32) float32 {
	centimeter = inches / 0.39370
	return centimeter
}

// InchestoMillimeters (inches float32) returns float32 in Millimeters value
func InchestoMillimeters(inches float32) float32 {
	millimeters = InchestoCentimeter(inches) * 10
	return millimeters
}

// InchestoFeet (inches float32) returns float32 in Feet value
func InchestoFeet(inches float32) float32 {
	feet = inches * 0.083333
	return feet
}

// MeterstoYard (meters float32) returns float32 in Yard value
func MeterstoYard(meters float32) float32 {
	yard = meters * yardmeter
	return yard
}

// YardtoMeters (yard float32) returns float32 in Meters value
func YardtoMeters(yard float32) float32 {
	meter = yard / yardmeter
	return meter
}

// YardtoFeet (feet float32) returns float32 in Yard value
func YardtoFeet(feet float32) float32 {
	feet = yard / 3.0000
	return feet
}
