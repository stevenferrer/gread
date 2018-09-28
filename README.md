[![Go Report Card](https://goreportcard.com/badge/github.com/steven-ferrer/gonsole)](https://goreportcard.com/report/github.com/steven-ferrer/gonsole) [![godocs](https://godoc.org/github.com/steven-ferrer/gonsole?status.svg)](https://godoc.org/github.com/steven-ferrer/gonsole) 
# gonsole
Golang package for reading inputs from stdin, file, and others that implements `io.Reader`

## Installing

`$ go get -v -u github.com/steven-ferrer/gonsole `

## Example

```go
package main

import (
    "fmt"
    "log"
    "math"
    "strings"

    "github.com/steven-ferrer/gonsole"
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
    
    // initialize a new instance of `gonsole.Reader`
    reader := gonsole.NewReader(strings.NewReader(s))
    //or if you want to read from stdin
    // reader := gonsole.NewReader(os.Stdin)
    
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
