package main

import "fmt"

type MyNum struct {
    x int
    y int
}

func Add(a, b MyNum) MyNum {
    return MyNum{a.x + b.x, a.y + b.y}
}

func Sub(a, b MyNum) MyNum {
    return MyNum{a.x - b.x, a.y - b.y}
}

func Mul(a, b MyNum) MyNum {
    var x = a.x*b.x - a.y*b.y
    var y = a.x*b.y + a.y*b.x
    return MyNum{x, y}
}

func main() {
    a, b := MyNum{1, 2}, MyNum{3, 1}
    c := Add(a, b)
    fmt.Println(c)
    d := Mul(a, b)
    fmt.Println(d)
}
