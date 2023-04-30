package main

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}
