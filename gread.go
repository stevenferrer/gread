package gread

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

//NextInt32 reads the next word and tries to convert it to int32
func (r *Reader) NextInt32() (int32, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	i32, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i32), nil
}

//NextUint32 reads the next word and tries to convert it to uint32
func (r *Reader) NextUint32() (uint32, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	ui32, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(ui32), nil
}

//NextInt64 reads the next word and tries to convert it to int64
func (r *Reader) NextInt64() (int64, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 64)
}

//NextUint64 reads the next word and tries to converts it to uint64
func (r *Reader) NextUint64() (uint64, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(s, 10, 64)
}

//NextFloat32 reads the next word and tries to convert it to float32
func (r *Reader) NextFloat32() (float32, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(f32), nil
}

//NextFloat64 reads the next word and tries to convert it to float64
func (r *Reader) NextFloat64() (float64, error) {
	s, err := r.NextWord()
	if err != nil {
		return 0, err
	}

	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f64, nil
}

//NextLine reads the next line
func (r *Reader) NextLine() (s string, err error) {
	hasNext := r.scanner.Scan()

	if !hasNext {
		return "", io.EOF
	}

	return r.scanner.Text(), r.scanner.Err()
}

//NextWord reads the next word
func (r *Reader) NextWord() (string, error) {
	if r.words.len() > 0 {
		word, err := r.words.pop()
		if err != nil {
			return "", err
		}
		return word, nil
	}

	//word stack is empty, get the next line
	line, err := r.NextLine()
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

//NewReader takes an io.Reader returns a new `gread.Reader`
func NewReader(rd io.Reader) *Reader {
	r := bufio.NewScanner(rd)
	return &Reader{r, &strStack{[]string{}}}
}

// NewReaderWithBufferSize takes an io reader and buffer size and a new `gread.Reader`
func NewReaderWithBufferSize(rd io.Reader, bufferSize int) *Reader {
	r := bufio.NewScanner(rd)
	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	r.Buffer(buf, bufferSize)
	return &Reader{r, &strStack{[]string{}}}
}
