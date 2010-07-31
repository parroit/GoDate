package GoDate

import "strconv"

type TimeSpan int64

const TimeSpan_MinValue = TimeSpan(-0x7FFFFFFFFFFFFFFF)
const TimeSpan_MaxValue = TimeSpan(0x7FFFFFFFFFFFFFFF)
const Zero = TimeSpan(0)
const TicksPerMillisecond int64 = 10000
const TicksPerSecond int64 = 10000 * 1000
const TicksPerMinute int64 = 10000 * 1000 * 60
const TicksPerHour int64 = 10000 * 1000 * 60 * 60
const TicksPerDay int64 = 10000 * 1000 * 60 * 60 * 24

func NewTimeSpan(days int64, hours int64, minutes int64, seconds int64, milliseconds int64) TimeSpan {
	value := days*TicksPerDay +
		hours*TicksPerHour +
		minutes*TicksPerMinute +
		seconds*TicksPerSecond +
		milliseconds*TicksPerMillisecond
	return TimeSpan(value)
}

func (t TimeSpan) Days() int {
	return int(int64(t) / TicksPerDay)
}
func (t TimeSpan) Hours() int {
	return int((int64(t) / TicksPerHour) % 24)
}
func (t TimeSpan) Minutes() int {
	return int((int64(t) / TicksPerMinute) % 60)
}
func (t TimeSpan) Seconds() int {
	return int((int64(t) / TicksPerSecond) % 60)
}

func (t TimeSpan) Milliseconds() int {
	return int((int64(t) / TicksPerMillisecond) % 1000)
}
func (t TimeSpan) Ticks() int64 {
	return int64(t)
}


func (t TimeSpan) TotalDays() float64 {
	return float64(t) / float64(TicksPerDay)
}
func (t TimeSpan) TotalHours() float64 {
	return float64(t) / float64(TicksPerHour)
}

func (t TimeSpan) TotalMinutes() float64 {
	return float64(t) / float64(TicksPerMinute)
}

func (t TimeSpan) TotalSeconds() float64 {
	return float64(t) / float64(TicksPerSecond)
}
func (t TimeSpan) TotalMilliseconds() float64 {
	return float64(t) / float64(TicksPerMillisecond)
}
func (t TimeSpan) String() string {
	days := t.Days()
	hours := t.Hours()
	minutes := t.Minutes()
	seconds := t.Seconds()
	fractional := t.Milliseconds()
	var result string
	if days != 0 {
		result = strconv.Itoa(days) + "."
	} else {
		result = ""
	}
	result = result + twoDigits(hours) + ":" +
		twoDigits(minutes) + ":" + twoDigits(seconds)

	if fractional != 0 {
		result = result + "." + strconv.Itoa(fractional)
	}
	return result
}


func twoDigits(value int) string {

	dec := strconv.Itoa(value / 10)
	uni := strconv.Itoa(value % 10)
	return dec + uni
}


// Convert a floating point number of days into a TimeSpan.
func FromDays(value float64) TimeSpan {
	return TimeSpan(value * float64(TicksPerDay))
}

// Convert a floating point number of hours into a TimeSpan.
func FromHours(value float64) TimeSpan {
	return TimeSpan(value * float64(TicksPerHour))
}

// Convert a floating point number of milliseconds into a TimeSpan.
func FromMilliseconds(value float64) TimeSpan {
	return TimeSpan(value * float64(TicksPerMillisecond))
}

// Convert a floating point number of minutes into a TimeSpan.
func FromMinutes(value float64) TimeSpan {
	return TimeSpan(value * float64(TicksPerMinute))
}


// Convert a floating point number of seconds into a TimeSpan.
func FromSeconds(value float64) TimeSpan {
	return TimeSpan(value * float64(TicksPerSecond))
}

