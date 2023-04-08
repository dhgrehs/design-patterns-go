// An example of Liskov Substitution Principle
package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rect struct {
	width, height int
}

func (r *Rect) GetWidth() int {
	return r.width
}

func (r *Rect) SetWidth(width int) {
	r.width = width
}

func (r *Rect) GetHeight() int {
	return r.height
}

func (r *Rect) SetHeight(height int) {
	r.height = height
}

// Liskov Substitution Principle:
// If a new type fulfills an interface (Sized) and works with a type T that implements the intf too,
// any further structure that aggregates that should also be usable.

type Square struct {
	Rect
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// SetWidth that protects width/height equivalence for squares, but will cause problems to DoIt function (and LS Principle).
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// Square2 that does not break the "inherited" behavior of a Rectangle that fulfilled Sized intf.
// Preserving LSP.
type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rect {
	return Rect{s.size, s.size}
}

func DoIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func main() {
	rc := &Rect{2, 3}
	DoIt(rc)

	sq := NewSquare(5)
	DoIt(sq)

	sq2 := Square2{5}
	rc2 := sq2.Rectangle()
	DoIt(&rc2)
}