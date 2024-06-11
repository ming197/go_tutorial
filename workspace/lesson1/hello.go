package main

import "fmt"

func main() {
    fmt.Println("Func Main!")
    test()
}

func init() {
    fmt.Println("Func Init!")
}

func test() {
     fmt.Println("Func Test!")
}