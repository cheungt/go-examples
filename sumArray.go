package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)

/* Takes 2 input from stdIn
input 1: N items to sum
input 2: list of N items seperated by ' '
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

    sum := 0

	// skip first scan
	scanner.Scan()
    scanner.Text()

	// list of ints
    scanner.Scan()
    line := scanner.Text()

    val := strings.Split(line, " ")

    for i := range val {
        n, _ := strconv.Atoi(val[i])
        sum += n
    }

    fmt.Println(sum)
}
