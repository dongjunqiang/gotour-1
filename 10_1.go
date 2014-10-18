package main

import "fmt"

func split(sum int) (x, y int) {
        //return sum * 4 / 9, sum - sum * 4/ 9
        x = sum * 4 / 9
        y = sum - x
        return y, x
}

func main() {
        fmt.Println(split(17))
}
