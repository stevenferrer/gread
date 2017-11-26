# gonsole
Golang package for reading inputs from console, file, and others that implements `io.Reader`

## Example


    package main

    import (
        "fmt"
        "log"
        "os"

        "github.com/steven-ferrer/gonsole"
    )

    func main() {
        reader := gonsole.NewReader(os.Stdin)
        fmt.Println("Enter int32: ")
        i32, err := reader.Int32()
        checkErr(err)
        fmt.Printf("%T is %d\n\n", i32, i32)

        fmt.Println("Enter int64: ")
        i64, err := reader.Int64()
        checkErr(err)
        fmt.Printf("%T is %d\n\n", i64, i64)

        fmt.Println("Enter uint32: ")
        ui32, err := reader.Uint32()
        checkErr(err)
        fmt.Printf("%T is %d\n\n", ui32, ui32)

        fmt.Println("Enter uint64: ")
        ui64, err := reader.Uint64()
        checkErr(err)
        fmt.Printf("%T is %d\n\n", ui64, ui64)

        fmt.Println("Enter float32: ")
        f32, err := reader.Float32()
        checkErr(err)
        fmt.Printf("%T is %f\n\n", f32, f32)

        fmt.Println("Enter float64: ")
        f64, err := reader.Float64()
        checkErr(err)
        fmt.Printf("%T is %f\n", f64, f64)
    }

    func checkErr(err error) {
        if err != nil {
            log.Fatal(err)
        }
    }

## Contributing

Please feel free to improve this by **openning an issue** or **submitting a PR**