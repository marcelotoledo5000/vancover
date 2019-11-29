package vancover

// KnottoKilomerh (knot float32) float32
func KnottoKilomerh(knot float32) float32 {
	return knot * 1.852
}

// KnottoMeterS (knot float32) float32
func KnottoMeterS(knot float32) float32 {
	return knot / 1.944
}

// KnottoFootS (knot float32) float32
func KnottoFootS(knot float32) float32 {
	return knot * 1.688
}

// KnottoMilesh (knot float32) float32
func KnottoMilesh(knot float32) float32 {
	return knot * 1.151
}
