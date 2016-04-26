package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
    scanner.Text()

    scanner.Scan()
    line := scanner.Text()

    val := strings.Split(line, " ")

    p, n, z := 0., 0., 0.
    for i := range val {
        v, _ := strconv.Atoi(val[i])
        switch {
            case v > 0:
            p++
            case v < 0:
            n++
            default:
            z++
        }
    }

    s := float64(len(val))

    fmt.Println(p/s)
    fmt.Println(n/s)
    fmt.Println(z/s)
}
