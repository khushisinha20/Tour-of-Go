package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1} //Y : 0 by default
    v3 = Vertex{} //X : 0 by default
    p = &Vertex{1, 2}
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
