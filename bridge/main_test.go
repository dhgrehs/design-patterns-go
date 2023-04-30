// Bridge is a structural design pattern that lets you split a large class or a set of closely related classes into
// two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

// In other terms it is a mechanism that decouples an interface from an implementation.
package main

import (
	"testing"
)

func main() {

}

// Problem to solve: Simplify the Circle drawing independently on what type of renderer I am using.
// Solution: Use the Bridge pattern, adding a Renderer interface in the Circle.
func TestBridge(t *testing.T) {
	rasterRenderer := RasterRenderer{}
	vectorRenderer := VectorRenderer{}

	circle := NewCircle(&vectorRenderer, 5)
	circle.Draw()

	circle = NewCircle(&rasterRenderer, 5)
	circle.Draw()
}
