package internal

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Time time.Time

func (t Time) ToString() (string, error) {
	return time.Time(t).UTC().Format(time.RFC3339Nano), nil
}

func (t *Time) FromString(s string) error {
	if dt, err := decodeTime(s); err != nil {
		return err
	} else {
		*t = Time(dt)
		return nil
	}
}

var reUnixtime = regexp.MustCompile(`^\d+(\.\d{1,9})?$`)

// decodeTime parses data bytes as time.Time in UTC timezone.
// Supported formats of the data bytes are:
// 1. RFC3339Nano string, e.g. "2006-01-02T15:04:05-07:00".
// 2. Date string, e.g. "2006-01-02".
// 3. Unix timestamp, e.g. "1136239445", "1136239445.8", "1136239445.812738".
func decodeTime(value string) (time.Time, error) {
	// Try parsing value as RFC3339 format.
	if t, err := time.ParseInLocation(time.RFC3339Nano, value, time.UTC); err == nil {
		return t.UTC(), nil
	}

	// Try parsing value as date format.
	if t, err := time.ParseInLocation("2006-01-02", value, time.UTC); err == nil {
		return t.UTC(), nil
	}

	// Try parsing value as timestamp, both integer and float formats supported.
	// e.g. "1618974933", "1618974933.284368".
	if reUnixtime.MatchString(value) {
		return decodeUnixtime(value)
	}

	return time.Time{}, errors.New("invalid time value")
}

// value must be valid unix timestamp, matches reUnixtime.
func decodeUnixtime(value string) (time.Time, error) {
	parts := strings.Split(value, ".")
	// Note: errors are ignored, since we already validated the value.
	sec, _ := strconv.ParseInt(parts[0], 10, 64)
	var nsec int64
	if len(parts) == 2 {
		nsec, _ = strconv.ParseInt(nanoSecondPrecision(parts[1]), 10, 64)
	}
	return time.Unix(sec, nsec).UTC(), nil
}

func nanoSecondPrecision(value string) string {
	return value + strings.Repeat("0", 9-len(value))
}
