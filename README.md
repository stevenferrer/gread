# gread

![Github Actions](https://github.com/sf9v/gread/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/gread/badge.svg?branch=master)](https://coveralls.io/github/sf9v/gread?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/gread)](https://goreportcard.com/report/github.com/sf9v/gread) [![godocs](https://godoc.org/github.com/sf9v/gread?status.svg)](https://godoc.org/github.com/sf9v/gread) 
[![Sourcegraph](https://sourcegraph.com/github.com/sf9v/gread/-/badge.svg)](https://sourcegraph.com/github.com/sf9v/gread?badge)


Go module for reading from anything that implements `io.Reader`

## Installing

```console
$ go get -v -u github.com/sf9v/gread
```

## Example

```go
package main

import (
    "fmt"
    "log"
    "math"
    "strings"

    "github.com/sf9v/gread"
)

func main() {
    s := fmt.Sprintf("%v %d %v %v %v %v\nanother line\nhello",
        math.MaxInt32,
        math.MaxInt64,
        math.MaxUint32,
        uint64(math.MaxUint64),
        math.MaxFloat32,
        math.MaxFloat64,
    )
    
    // initialize a new instance of `gread.Reader`
    reader := gread.NewReader(strings.NewReader(s))
    //or if you want to read from stdin
    // reader := gread.NewReader(os.Stdin)
    
    i32, err := reader.NextInt32()
    checkErr(err)
    fmt.Printf("%T is %d\n\n", i32, i32)

    i64, err := reader.NextInt64()
    checkErr(err)
    fmt.Printf("%T is %d\n\n", i64, i64)

    ui32, err := reader.NextUint32()
    checkErr(err)
    fmt.Printf("%T is %d\n\n", ui32, ui32)

    ui64, err := reader.NextUint64()
    checkErr(err)
    fmt.Printf("%T is %d\n\n", ui64, ui64)

    f32, err := reader.NextFloat32()
    checkErr(err)
    fmt.Printf("%T is %f\n\n", f32, f32)

    f64, err := reader.NextFloat64()
    checkErr(err)
    fmt.Printf("%T is %f\n", f64, f64)

    w, err := reader.NextWord()
    checkErr(err)
    fmt.Printf("\nword1 is %q\n", w)
    w, err = reader.NextWord()
    checkErr(err)
    fmt.Printf("word2 is %q\n", w)

    l, err := reader.NextLine()
    checkErr(err)
    fmt.Printf("\nline is %q\n", l)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
```

## Contributing

Please feel free to improve this by **openning an issue** or **submitting a PR**
