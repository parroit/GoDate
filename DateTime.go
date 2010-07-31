/*
	GoDate - A Go package for date  and time manipulation
    Copyright (C) 2010 Andrea Parodi

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package GoDate

import "time"

// Days in each month in regular years.
var daysForEachMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// Days before each month in regular years.
var daysBeforeMonth = []int{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30}


// The number of days in 400 years.
const DaysPer400Years = (400*365 + 97)

// The number of days in 100 years, when the 100th is not a leap year.
const DaysPer100Years = (100*365 + 24)

// The number of days in 4 years, including the leap day.
const DaysPer4Years = (4*365 + 1)


type DateTime uint64


const MinValue = DateTime(0)
const MaxValue = DateTime(0x8000000000000000)

// Determine if a year is a leap year.
func IsLeapYear(year int) bool {
	return ((year%4) == 0 &&
		((year%100) != 0 || (year%400) == 0))
}

// Convert a year into a number of ticks.
func yearToTicks(year int) DateTime {
	year--
	return DateTime(int64(year*365+year/4-
		year/100+year/400) * TicksPerDay)
}


func NewDate(year int, month int, day int) DateTime {
	return dateToTicks(year, month, day)
}


// Convert a YMD date into a number of ticks.
func dateToTicks(year int, month int, day int) DateTime {
	var daysInMonth int
	var result int64
	var isLeap bool
	if year >= 1 && year <= 9999 &&
		month >= 1 && month <= 12 {

		isLeap = IsLeapYear(year)
		daysInMonth = daysForEachMonth[month-1]
		if month == 2 && isLeap {
			daysInMonth++
		}
		if day >= 1 && day <= daysInMonth {

			result = (int64(yearToTicks(year)) / TicksPerDay)
			result += int64(daysBeforeMonth[month-1])
			if month > 2 && isLeap {
				result++
			}
			return DateTime((result + int64(day-1)) * TicksPerDay)
		}
	}
	return MinValue
}

func (d DateTime) Day() int {
	// Get the current year.
	year := d.Year()

	// Get the tick offset within the year for the day.
	ticks := d - yearToTicks(year)

	// Convert the tick offset into days.
	days := int(int64(ticks) / TicksPerDay)

	// Adjust for Jan and Feb in leap years.
	if IsLeapYear(year) {
		if days < 31 {
			return days + 1
		} else if days < (31 + 29) {
			return days - 30
		}
		days--
	}

	// Search for the starting month.
	month := 1
	for month < 12 && days >= daysBeforeMonth[month] {
		month++
	}
	return days - daysBeforeMonth[month-1] + 1
}


func (d DateTime) Month() int {
	// Get the current year.
	year := d.Year()

	// Get the tick offset within the year for the day.
	ticks := d - yearToTicks(year)

	// Convert the tick offset into days.
	days := int(int64(ticks) / TicksPerDay)

	// Adjust for Jan and Feb in leap years.
	if IsLeapYear(year) {
		if days < 31 {
			return 1
		} else if days < (31 + 29) {
			return 2
		}
		days--
	}

	// Search for the starting month.
	month := 1
	for month < 12 && days >= daysBeforeMonth[month] {
		month++
	}

	return month
}

func (d DateTime) Year() int {
	// Note: there is probably a tricky mathematical
	// formula for doing this, but this version is a
	// lot easier to understand and debug.

	// Convert the tick count into a day value.
	days := int(int64(d) / TicksPerDay)

	// Determine the 400-year cycle that contains the date.
	yearBase := ((days / DaysPer400Years) * 400) + 1
	yearOffset := days % DaysPer400Years

	// Determine the 100-year cycle that contains the date.
	if yearOffset >= DaysPer100Years*3 {
		yearBase += 300
		yearOffset -= DaysPer100Years * 3
		if yearOffset >= DaysPer100Years {
			// The date is at the end of a 400-year cycle.
			return yearBase + 399
		}
	} else if yearOffset >= DaysPer100Years*2 {
		yearBase += 200
		yearOffset -= DaysPer100Years * 2
	} else if yearOffset >= DaysPer100Years {
		yearBase += 100
		yearOffset -= DaysPer100Years
	}

	// Determine the 4-year cycle that contains the date.
	temp := yearOffset / DaysPer4Years
	yearBase += temp * 4
	yearOffset -= temp * DaysPer4Years

	// Determine the year out of the 4-year cycle.
	if yearOffset >= 365*3 {
		return yearBase + 3
	} else if yearOffset >= 365*2 {
		return yearBase + 2
	} else if yearOffset >= 365 {
		return yearBase + 1
	} else {
		return yearBase
	}

	return 0
}

func NewDateTime(year int, month int, day int, hour int, minute int, second int) DateTime {
	return dateToTicks(year, month, day) + timeToTicks(hour, minute, second)
}
func NewTime(hour int, minute int, second int) DateTime {
	return timeToTicks(hour, minute, second)
}
// Convert a time into a number of ticks.
func timeToTicks(hour int, minute int, second int) DateTime {
	if hour >= 0 && hour <= 23 &&
		minute >= 0 && minute <= 59 &&
		second >= 0 && minute <= 59 {
		return DateTime(int64(hour*3600+minute*60+second) *
			TicksPerSecond)
	}
	return MinValue
}

func (d DateTime) Date() DateTime {
	return d - (d % DateTime(TicksPerDay))
}


func (d DateTime) Hour() int {
	//					return (int)((value_ / (ticksPerDay / 24)) % 24);

	return int((d / DateTime(TicksPerDay/24)) % 24)
}

func (d DateTime) Minute() int {
	return int((d / (DateTime(TicksPerDay) / (24 * 60))) % 60)
}

func (d DateTime) Second() int {
	return int((d / 10000000) % 60)
}

func (d DateTime) Millisecond() int {
	return int((d / 10000) % 1000)
}

func (d DateTime) TimeOfDay() TimeSpan {
	return TimeSpan(d % DateTime(TicksPerDay))
}

func FromGoTime(t time.Time) DateTime {
	return NewDateTime(int(t.Year), int(t.Month), int(t.Day),
		int(t.Hour), int(t.Minute), int(t.Second))
}

func Now() DateTime {
	return FromGoTime(*time.LocalTime())
}

func UtcNow() DateTime {
	return FromGoTime(*time.UTC())
}

func Today() DateTime {
	return FromGoTime(*time.LocalTime()).Date()
}


// Determine the number of days in a month.
func DaysInMonth(year int, month int) int {

	if IsLeapYear(year) && month == 2 {
		return 29
	} else {
		return daysForEachMonth[month-1]
	}
	return 0
}
func (d DateTime) DayOfWeek() int {
	return int(((d / DateTime(TicksPerDay)) + 1) % 7)
}

func (d DateTime) DayOfYear() int {
	// Get the tick offset within the year for the day.
	ticks := d - yearToTicks(d.Year())

	// Convert the tick offset into days.
	return int(ticks/DateTime(TicksPerDay)) + 1
}


// Add certain periods of time to this DateTime value.
func (d DateTime) AddDays(value float64) DateTime {
	return d + DateTime(FromDays(value))
}
func (d DateTime) AddHours(value float64) DateTime {
	return d + DateTime(FromHours(value))
}
func (d DateTime) AddMilliseconds(value float64) DateTime {
	return d + DateTime(FromMilliseconds(value))
}
func (d DateTime) AddMinutes(value float64) DateTime {
	return d + DateTime(FromMinutes(value))
}
func (d DateTime) AddSeconds(value float64) DateTime {
	return d + DateTime(FromSeconds(value))
}
func (d DateTime) Format(layout string) string{
	return d.ToGoTime().Format(layout)
}

func (d DateTime) String() string {
	return d.ToGoTime().String()
}

func (d DateTime) ToGoTime() *time.Time {
	return &time.Time{Year:int64(d.Year()),Month:d.Month(),Day:d.Day(),
		Hour:d.Hour(),Minute:d.Minute(),Second:d.Second()}
	
}


