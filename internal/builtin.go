package internal

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type String string

func (sv String) ToString() (string, error) {
	return string(sv), nil
}

func (sv *String) FromString(s string) error {
	*sv = String(s)
	return nil
}

type Bool bool

func (bv Bool) ToString() (string, error) {
	return strconv.FormatBool(bool(bv)), nil
}

func (bv *Bool) FromString(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*bv = Bool(v)
	return nil
}

type Int int

func (iv Int) ToString() (string, error) {
	return strconv.Itoa(int(iv)), nil
}

func (iv *Int) FromString(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*iv = Int(v)
	return nil
}

type Int8 int8

func (iv Int8) ToString() (string, error) {
	return strconv.FormatInt(int64(iv), 10), nil
}

func (iv *Int8) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return err
	}
	*iv = Int8(v)
	return nil
}

type Int16 int16

func (iv Int16) ToString() (string, error) {
	return strconv.FormatInt(int64(iv), 10), nil
}

func (iv *Int16) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return err
	}
	*iv = Int16(v)
	return nil
}

type Int32 int32

func (iv Int32) ToString() (string, error) {
	return strconv.FormatInt(int64(iv), 10), nil
}

func (iv *Int32) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*iv = Int32(v)
	return nil
}

type Int64 int64

func (iv Int64) ToString() (string, error) {
	return strconv.FormatInt(int64(iv), 10), nil
}

func (iv *Int64) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*iv = Int64(v)
	return nil
}

type Uint uint

func (uv Uint) ToString() (string, error) {
	return strconv.FormatUint(uint64(uv), 10), nil
}

func (uv *Uint) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*uv = Uint(v)
	return nil
}

type Uint8 uint8

func (uv Uint8) ToString() (string, error) {
	return strconv.FormatUint(uint64(uv), 10), nil
}

func (uv *Uint8) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return err
	}
	*uv = Uint8(v)
	return nil
}

type Uint16 uint16

func (uv Uint16) ToString() (string, error) {
	return strconv.FormatUint(uint64(uv), 10), nil
}

func (uv *Uint16) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return err
	}
	*uv = Uint16(v)
	return nil
}

type Uint32 uint32

func (uv Uint32) ToString() (string, error) {
	return strconv.FormatUint(uint64(uv), 10), nil
}

func (uv *Uint32) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	*uv = Uint32(v)
	return nil
}

type Uint64 uint64

func (uv Uint64) ToString() (string, error) {
	return strconv.FormatUint(uint64(uv), 10), nil
}

func (uv *Uint64) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*uv = Uint64(v)
	return nil
}

type Float32 float32

func (fv Float32) ToString() (string, error) {
	return strconv.FormatFloat(float64(fv), 'f', -1, 32), nil
}

func (fv *Float32) FromString(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*fv = Float32(v)
	return nil
}

type Float64 float64

func (fv Float64) ToString() (string, error) {
	return strconv.FormatFloat(float64(fv), 'f', -1, 64), nil
}

func (fv *Float64) FromString(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*fv = Float64(v)
	return nil
}

type Complex64 complex64

func (cv Complex64) ToString() (string, error) {
	return strconv.FormatComplex(complex128(cv), 'f', -1, 64), nil
}

func (cv *Complex64) FromString(s string) error {
	v, err := strconv.ParseComplex(s, 64)
	if err != nil {
		return err
	}
	*cv = Complex64(v)
	return nil
}

type Complex128 complex128

func (cv Complex128) ToString() (string, error) {
	return strconv.FormatComplex(complex128(cv), 'f', -1, 128), nil
}

func (cv *Complex128) FromString(s string) error {
	v, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return err
	}
	*cv = Complex128(v)
	return nil
}

type Time time.Time

func (tv Time) ToString() (string, error) {
	return time.Time(tv).UTC().Format(time.RFC3339Nano), nil
}

func (tv *Time) FromString(s string) error {
	if t, err := DecodeTime(s); err != nil {
		return err
	} else {
		*tv = Time(t)
		return nil
	}
}

var reUnixtime = regexp.MustCompile(`^\d+(\.\d{1,9})?$`)

// DecodeTime parses data bytes as time.Time in UTC timezone.
// Supported formats of the data bytes are:
// 1. RFC3339Nano string, e.g. "2006-01-02T15:04:05-07:00".
// 2. Date string, e.g. "2006-01-02".
// 3. Unix timestamp, e.g. "1136239445", "1136239445.8", "1136239445.812738".
func DecodeTime(value string) (time.Time, error) {
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
		return DecodeUnixtime(value)
	}

	return time.Time{}, errors.New("invalid time value")
}

// value must be valid unix timestamp, matches reUnixtime.
func DecodeUnixtime(value string) (time.Time, error) {
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

// ByteSlice is a wrapper of []byte to implement StringConverter.
// NOTE: we're using base64.StdEncoding here, not base64.URLEncoding.
type ByteSlice []byte

func (bs ByteSlice) ToString() (string, error) {
	return base64.StdEncoding.EncodeToString(bs), nil
}

func (bs *ByteSlice) FromString(s string) error {
	v, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*bs = ByteSlice(v)
	return nil
}
