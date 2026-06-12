package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Describe() string
}

type CloneableShape interface {
	Shape
	Clone() Shape
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (c *Circle) Describe() string { return fmt.Sprintf("Circle(r=%.1f)", c.Radius) }
func (c *Circle) Clone() Shape     { return &Circle{Radius: c.Radius} }

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64    { return r.Width * r.Height }
func (r *Rectangle) Describe() string { return fmt.Sprintf("Rect(%.1fx%.1f)", r.Width, r.Height) }
func (r *Rectangle) Clone() Shape     { return &Rectangle{Width: r.Width, Height: r.Height} }

// Point implements Shape but not CloneableShape
type Point struct {
	X, Y float64
}

func (p *Point) Area() float64    { return 0 }
func (p *Point) Describe() string { return fmt.Sprintf("Point(%.1f, %.1f)", p.X, p.Y) }

func cloneAll(shapes []Shape) ([]Shape, []error) {
	copies := make([]Shape, len(shapes))
	errs := make([]error, len(shapes))

	for i, s := range shapes {
		c, ok := s.(CloneableShape)
		if !ok {
			errs[i] = fmt.Errorf("%T does not implement Clone", s)
			copies[i] = s
			continue
		}
		copies[i] = c.Clone()
	}
	return copies, errs
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Width: 10, Height: 3},
		&Point{X: 1, Y: 2},
	}

	copies, errs := cloneAll(shapes)

	copies[0].(*Circle).Radius = 99

	fmt.Println("Originals:")
	for _, s := range shapes {
		fmt.Printf("  %s\n", s.Describe())
	}
	fmt.Println("Copies:")
	for _, s := range copies {
		fmt.Printf("  %s\n", s.Describe())
	}
	fmt.Println("Errors:")
	for _, err := range errs {
		if err != nil {
			fmt.Printf("  %v\n", err)
		}
	}
}
