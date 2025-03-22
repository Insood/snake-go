package main

import raylib "github.com/gen2brain/raylib-go/raylib"

type Position struct {
	X float32
	Y float32
}

type Movement struct {
	X     float32
	Y     float32
	Speed int64
}

type Color struct {
	Color raylib.Color
}

type Head struct {
}

type Food struct {
}

type Segment struct {
	NextSegmentID ID
	LastX         float32
	LastY         float32
}

func NewPosition(x, y float32) *Position {
	return &Position{X: x, Y: y}
}
func NewMovement(speed int64, x, y float32) *Movement {
	return &Movement{X: x, Y: y, Speed: speed}
}

func NewColor(color raylib.Color) *Color {
	return &Color{Color: color}
}

func NewHead() *Head {
	return &Head{}
}

func NewFood() *Food {
	return &Food{}
}

func NewSegment() *Segment {
	return &Segment{NextSegmentID: -1}
}
