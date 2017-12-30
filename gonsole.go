package gonsole

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type strStack struct{ arr []string }

func (st *strStack) len() int { return len(st.arr) }

func (st *strStack) clear() { st.arr = []string{} }

func (st *strStack) push(s string) { st.arr = append(st.arr, s) }

func (st *strStack) pushArray(s ...string) { st.arr = append(st.arr, s...) }

func (st *strStack) pop() (string, error) {
	if len(st.arr) > 0 {
		var s string
		s, st.arr = st.arr[0], st.arr[1:]

		return s, nil
	}

	return "", errors.New("stringStack: string stack is empty")
}

//Reader is a wrapper for *bufio.Scanner
type Reader struct {
	scanner *bufio.Scanner

	//words will contain the words in a line
	words *strStack
}

//Int32 reads the next word and tries to convert it to int32
func (r *Reader) Int32() (int32, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	i32, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i32), nil
}

//Uint32 reads the next word and tries to convert it to uint32
func (r *Reader) Uint32() (uint32, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	ui32, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(ui32), nil
}

//Int64 reads the next word and tries to convert it to int64
func (r *Reader) Int64() (int64, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 64)
}

//Uint64 reads the next word and tries to converts it to uint64
func (r *Reader) Uint64() (uint64, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(s, 10, 64)
}

//Float32 reads the next word and tries to convert it to float32
func (r *Reader) Float32() (float32, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(f32), nil
}

//Float64 reads the next word and tries to convert it to float64
func (r *Reader) Float64() (float64, error) {
	s, err := r.Word()
	if err != nil {
		return 0, err
	}

	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f64, nil
}

//Line reads the next line
func (r *Reader) Line() (s string, err error) {
	if r.scanner.Scan() {
		s = r.scanner.Text()
	}
	err = r.scanner.Err()
	return
}

//Word reads the next word
func (r *Reader) Word() (string, error) {
	if r.words.len() > 0 {
		word, err := r.words.pop()
		if err != nil {
			return "", err
		}
		return word, nil
	}

	//word stack is empty, we can a new line
	line, err := r.Line()
	if err != nil {
		return "", err
	}

	words := strings.Fields(line)
	if len(words) > 0 {
		r.words.clear()
		r.words.pushArray(words...)

		word, _ := r.words.pop()
		return word, nil
	}

	return "", nil
}

//NewReader takes an io.Reader returns a new `gonsole.Reader`
func NewReader(rd io.Reader) *Reader {
	r := bufio.NewScanner(rd)
	return &Reader{r, &strStack{[]string{}}}
}
