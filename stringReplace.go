package main
import (
    "fmt"
    "strconv"
)


/*
Takes input of time hh:mm:ss(PM|AM)
prints out time in 24h
*/
func main() {
    var time string
    fmt.Scanf("%s\n", &time)

    am := time[len(time) - 2] == 'A'
    h, _ := strconv.Atoi(time[:2])
    if !am && h != 12{
        h += 12
        time = strconv.Itoa(h) + time[2:]
    } else if am && h == 12 {
        time = "00" + time[2:]
    }

    fmt.Println(time[0:len(time)-2])
}
