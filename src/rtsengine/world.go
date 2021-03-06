package rtsengine

import (
	"image"
	"math/rand"
	"time"
)

/*
 World 2D grid. That is an array of acre structures.

*/

// World maintains the world state. This is the big one!
type World struct {
	Grid
}

// NewWorld will construct a random world of width and height specified.
// works on 'this'. Another way of thinking is width are the columns
// and height are the rows.
func NewWorld(width int, height int) *World {
	world := World{}

	// When the worldLocation is 0,0 then the grid IS the world.
	world.GenerateGrid(image.Point{0, 0}, width, height)

	return &world
}

// GenerateSimple will generate a simple world for basic testing.
// Good for testing pathing etcetera.
func (world *World) GenerateSimple() {

	// Make all the world grass!
	for i := range world.Matrix {
		for j := range world.Matrix[i] {
			world.Matrix[i][j].unit = nil
			world.Matrix[i][j].terrain = Grass
		}
	}

	// Randomly dot with trees and mountains
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Trees
	for i := 0; i < 10000; i++ {
		xr := r1.Intn(world.Span.Dx())
		yr := r1.Intn(world.Span.Dy())
		world.Matrix[xr][yr].terrain = Trees
	}

	// Mountains
	for i := 0; i < 10000; i++ {
		xr := r1.Intn(world.Span.Dx())
		yr := r1.Intn(world.Span.Dy())
		world.Matrix[xr][yr].terrain = Mountains
	}

	centerPoint := world.Center()
	world.Matrix[centerPoint.X][centerPoint.Y].terrain = Grass

	// Generate a straight fence and go through anything.
	points := world.DirectLineBresenham(&image.Point{20, 20}, &image.Point{40, 40})

	for _, point := range points {
		world.Matrix[point.X][point.Y].terrain = Mountains
	}
}

// GenerateGrassWorld will generate a world of grass....
func (world *World) GenerateGrassWorld() {

	// Make all the world grass!
	for i := range world.Matrix {
		for j := range world.Matrix[i] {
			world.Matrix[i][j].unit = nil
			world.Matrix[i][j].terrain = Grass
		}
	}

	centerPoint := world.Center()
	world.Matrix[centerPoint.X][centerPoint.Y].terrain = Grass

}

// Center returns the x,y center of this View.
func (world *World) Center() image.Point {
	return image.Point{world.Span.Min.X + (world.Span.Dx() / 2), world.Span.Min.Y + (world.Span.Dy() / 2)}
}
