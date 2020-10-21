// cf converts its numeric argument to celcius or farenhiet
package main

import (
    "fmt"
    "os"
    "strconv"
    "github.com/makarkul/gopl/ch2/tempconv"
)

func main() {

    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := tempconv.Farenheit(t)
        c := tempconv.Celcius(t)

        fmt.Printf("%s = %s, %s = %s\n",
            f, tempconv.FToC(f), c, tempconv.CToF(c))
    }
}
