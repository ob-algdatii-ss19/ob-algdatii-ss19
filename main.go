package main

import (
	"fmt"
	"github.com/ob-algdatii-ss19/ob-algdatii-ss19/graph"
)

func main() {
	g := graph.NewGraphAdjMat(10)
	g.AddEdge(1, 2)
	print(g)
}

func print(g graph.Graph) {
	fmt.Printf("edges = %v\n", g.Edges())
}
