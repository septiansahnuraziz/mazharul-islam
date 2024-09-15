package utils

import "time"

// ParseDurationWithDefault parses a duration string and returns the parsed duration.
// If the parsing fails, it returns the default duration provided.
func ParseDurationWithDefault(input string, defaultDuration time.Duration) time.Duration {
	parsedDuration, err := time.ParseDuration(input)
	if err != nil {
		return defaultDuration
	}

	return parsedDuration
}

// ParseDate parses a date string using the specified layout and returns the corresponding time.Time value.
// If the parsing fails, it returns the zero time value.
func ParseDate(layout, value string) time.Time {
	date, _ := time.Parse(layout, value)

	return date
}

// ParseDatetimeToRFC3339 formats the provided time value to the RFC3339 format.
func ParseDatetimeToRFC3339(inputTime *time.Time) string {
	return inputTime.Format(time.RFC3339)
}

func GetNowTime() time.Time {
	return time.Now()
}

func InMinuteTimeRange(startTime time.Time, stopTime uint) bool {
	return time.Since(startTime) >= time.Duration(stopTime)*time.Minute
}

func GetNowTimeRFC3339() string {
	times := time.Now()

	return times.Format(time.RFC3339)
}

func GetDate(times time.Time) time.Time {
	date := time.Date(times.Year(), times.Month(), times.Day(), 0, 0, 0, 0, time.UTC)

	return date
}

func GetTimeDuration(param int) time.Duration {
	return time.Duration(param)
}

func AddTime(currentTime time.Time, additionalTime int, unit string) time.Time {
	var timeUnit time.Duration

	switch unit {
	case "second":
		timeUnit = time.Second
	case "minute":
		timeUnit = time.Minute
	case "hour":
		timeUnit = time.Hour
	}

	return currentTime.Add(time.Duration(additionalTime) * timeUnit)
}

func SubTime(currentTime time.Time, subtractionTime int, unit string) time.Time {
	var timeUnit time.Duration

	switch unit {
	case "second":
		timeUnit = time.Second
	case "minute":
		timeUnit = time.Minute
	case "hour":
		timeUnit = time.Hour
	}

	return currentTime.Add(time.Duration(-subtractionTime) * timeUnit)
}

func GetTomorrowDate(currentTime time.Time) time.Time {
	yyyy, mm, dd := currentTime.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, currentTime.Location())
	return tomorrow
}
