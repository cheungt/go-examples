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
}
