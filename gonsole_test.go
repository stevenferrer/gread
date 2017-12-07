package gonsole

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

//TODO: Test for reading from file

func TestInt32InRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt32, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.Int32()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatInt(math.MinInt32, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Int32()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestInt32OutOfRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt32+1, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.Int32()
	if err == nil {
		t.Error("max: expecting non-nil error got ", err)
	}

	s = strconv.FormatInt(math.MinInt32-1, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Int32()
	if err == nil {
		t.Error("min: expecting non-nil error got ", err)
	}
}

func TestUint32InRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint32, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.Uint32()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatUint(0, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Uint32()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestUint32OutOfRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint32+1, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.Uint32()
	if err == nil {
		t.Error("Expecting non-nil error got ", err)
	}
}

func TestInt64InRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt64, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.Int64()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatInt(math.MinInt64, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Int64()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestUint64InRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint64, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.Uint64()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatUint(0, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Uint64()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestFloat32InRange(t *testing.T) {
	s := strconv.FormatFloat(math.MaxFloat32, 'E', -1, 32)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.Float32()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}

	s = strconv.FormatFloat(math.SmallestNonzeroFloat32, 'E', -1, 32)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Float32()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}
}

func TestFloat64InRange(t *testing.T) {
	s := strconv.FormatFloat(math.MaxFloat64, 'E', -1, 64)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.Float64()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}

	s = strconv.FormatFloat(math.SmallestNonzeroFloat64, 'E', -1, 64)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.Float64()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}
}

func TestLine(t *testing.T) {
	s := "The quick brown fox jumps over the lazy dog.\n"
	reader := NewReader(strings.NewReader(s))
	s2, err := reader.Line()
	if err != nil {
		t.Error("expecting nil error got", err)
	}

	if strings.TrimSpace(s) != s2 {
		t.Errorf("%q is not equal to %q", s, s2)
	}
}
