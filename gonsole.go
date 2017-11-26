package gonsole

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

//Reader is a wrapper for *bufio.Reader
type Reader struct {
	reader *bufio.Reader
}

//Int32 reads the next line and tries
//to convert it to int32
func (r *Reader) Int32() (int32, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	i32, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i32), nil
}

//Uint32 reads the next line and tries
//to convert it to uint32
func (r *Reader) Uint32() (uint32, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	ui32, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(ui32), nil
}

//Int64 reads the next line and tries
//to convert it to int64
func (r *Reader) Int64() (int64, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i64, nil
}

//Uint64 reads the next line and tries
//to converts it to uint64
func (r *Reader) Uint64() (uint64, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return ui64, nil
}

//Float32 reads the next line and tries
//to convert it to float32
func (r *Reader) Float32() (float32, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(f32), nil
}

//Float64 reads the next line and tries
//to convert it to float64
func (r *Reader) Float64() (float64, error) {
	s, err := r.Line()
	if err != nil {
		return 0, err
	}

	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f64, nil
}

//Line reads the next line (terminated by newline) and trims
//the extra spaces around including newline
func (r *Reader) Line() (string, error) {
	s, err := r.reader.ReadString('\n')
	//remove spaces and extra lines
	s = strings.TrimSpace(s)
	return s, err
}

//NewReader takes an io.Reader returns a new console reader
func NewReader(rd io.Reader) *Reader {
	r := bufio.NewReader(rd)
	return &Reader{r}
}
