package vancover

const (
	secondhour = 0.00027778
	seconday   = 0.000011574
)

var (
	second  float32
	hour    float32
	day     float32
	minutes float32
	week    float32
)

// SecondtoHour (second float32) float32 in hour value
func SecondtoHour(second float32) float32 {
	hour = second * secondhour
	return hour
}

// SecondtoDay (second float32) float32 in hour value
func SecondtoDay(second float32) float32 {
	day = second * seconday
	return day
}

// HourtoSecond (hour float32) float32 in second value
func HourtoSecond(hour float32) float32 {
	second = hour / secondhour
	return second
}

// HourtoMinutes (hour float32) float32 in minutes value
func HourtoMinutes(hour float32) float32 {
	minutes = hour * 60
	return minutes
}

// DaytoSecond (day float32) float32 in second value
func DaytoSecond(day float32) float32 {
	second = day / seconday
	return second
}

// DaytoMinutes (day float32) float32 in minutes value
func DaytoMinutes(day float32) float32 {
	minutes = day * 1440
	return minutes
}

// MinutestoDay (day float32) float32 in minutes value
func MinutestoDay(minutes float32) float32 {
	day = minutes * 0.00069444
	return day
}

// MinutestoHour (minutes float32) float32 in hour value
func MinutestoHour(minutes float32) float32 {
	hour = minutes / 60
	return hour
}

// MinutestoSecond (minutes float32) float32 in second value
func MinutestoSecond(minutes float32) float32 {
	second = minutes * 60
	return second
}

// MinutestoWeek (minutes float32) float32 in week value
func MinutestoWeek(minutes float32) float32 {
	week = MinutestoDay(minutes) * 7
	return week
}

// WeektoDays (minutes float32) float32 in days value
func WeektoDays(week float32) float32 {
	day = week / 7
	return day
}
