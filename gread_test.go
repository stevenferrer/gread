package gread

import (
	"bufio"
	"math"
	"strconv"
	"strings"
	"testing"
)

//TODO: Test for reading from file

func TestNextInt32InRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt32, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.NextInt32()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatInt(math.MinInt32, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextInt32()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestNextInt32OutOfRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt32+1, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.NextInt32()
	if err == nil {
		t.Error("max: expecting non-nil error got ", err)
	}

	s = strconv.FormatInt(math.MinInt32-1, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextInt32()
	if err == nil {
		t.Error("min: expecting non-nil error got ", err)
	}
}

func TestNextUint32InRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint32, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.NextUint32()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatUint(0, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextUint32()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestNextUint32OutOfRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint32+1, 10)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.NextUint32()
	if err == nil {
		t.Error("Expecting non-nil error got ", err)
	}
}

func TestNextInt64InRange(t *testing.T) {
	s := strconv.FormatInt(math.MaxInt64, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.NextInt64()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatInt(math.MinInt64, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextInt64()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestNextUint64InRange(t *testing.T) {
	s := strconv.FormatUint(math.MaxUint64, 10)
	reader := NewReader(strings.NewReader(s))

	_, err := reader.NextUint64()
	if err != nil {
		t.Error("max: expecting nil error got", err)
	}

	s = strconv.FormatUint(0, 10)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextUint64()
	if err != nil {
		t.Error("min: expecting nil error got", err)
	}
}

func TestNextFloat32InRange(t *testing.T) {
	s := strconv.FormatFloat(math.MaxFloat32, 'E', -1, 32)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.NextFloat32()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}

	s = strconv.FormatFloat(math.SmallestNonzeroFloat32, 'E', -1, 32)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextFloat32()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}
}

func TestNextFloat64InRange(t *testing.T) {
	s := strconv.FormatFloat(math.MaxFloat64, 'E', -1, 64)
	reader := NewReader(strings.NewReader(s))
	_, err := reader.NextFloat64()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}

	s = strconv.FormatFloat(math.SmallestNonzeroFloat64, 'E', -1, 64)
	reader = NewReader(strings.NewReader(s))
	_, err = reader.NextFloat64()
	if err != nil {
		t.Error("expecting a nil error got", err)
	}
}

func TestNextLine(t *testing.T) {
	s := "The quick brown fox jumps over the lazy dog.\n"
	reader := NewReader(strings.NewReader(s))
	s2, err := reader.NextLine()
	if err != nil {
		t.Error("expecting nil error got", err)
	}

	if strings.TrimSpace(s) != s2 {
		t.Errorf("%q is not equal to %q", s, s2)
	}
}

func TestStringStack(t *testing.T) {
	words := &strStack{strings.Fields("the quick")}

	s, _ := words.pop()
	if s != "the" {
		t.Errorf("expecting \"the\" got %s", s)
	}

	s, _ = words.pop()
	if s != "quick" {
		t.Errorf("expecting \"quick\" got %s", s)
	}

	_, err := words.pop()
	if err == nil {
		t.Errorf("expecting non-nil error")
	}
}

func TestNextWord(t *testing.T) {
	str := "first line\nsecond line erf\nthird line\n"
	lines := strings.Split(str, "\n")

	reader := NewReader(strings.NewReader(str))

	s, err := reader.NextLine()
	if err != nil {
		t.Error("expecting nil error got", err)
	}
	if s != lines[0] {
		t.Errorf("expecting %s got %q", lines[0], s)
	}

	s, err = reader.NextWord()
	if err != nil {
		t.Error("expecting nil error got", err)
	}

	if s != "second" {
		t.Errorf("expecting \"second\" got %q", s)
	}

	s, err = reader.NextWord()
	if err != nil {
		t.Error("expecting nil error got", err)
	}

	if s != "line" {
		t.Errorf("expecting \"line\" got %q", s)
	}

	s, err = reader.NextLine()
	if err != nil {
		t.Error("expecting nil error got", err)
	}
	if s != lines[2] {
		t.Errorf("expecting %s got %q", lines[2], s)
	}

}

func TestLongLinesThatExceedsDefaultBuffer(t *testing.T) {

	strs := []string{}

	for i := 0; i < bufio.MaxScanTokenSize+64; i++ {
		strs = append(strs, "abcd")
	}

	b := strings.NewReader(strings.Join(strs, ""))
	reader := NewReader(b)

	_, err := reader.NextLine()
	if err == nil {
		t.Error("expecting error but got", err)
	}

	bufferSize := 1024 * 1024
	b = strings.NewReader(strings.Join(strs, ""))
	reader = NewReaderWithBufferSize(b, bufferSize)

	_, err = reader.NextLine()
	if err != nil {
		t.Error("expecting nil error but got", err)
	}
}
