package main

import (
    "fmt"
    "mycalc/mynum"
)

func main() {
    a, b := mynum.MyNum{1, 2}, mynum.MyNum{3, 1}
    c := mynum.Add(a, b)
    fmt.Println(c)
    d := mynum.Mul(a, b)
    fmt.Println(d)
}
