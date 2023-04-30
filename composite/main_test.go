// Composite is a partitioning design pattern.
// It is a mechanism for treating a group of objects the same way as a single instance of the object.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A GraphicObject can be single or be a group of GraphicObjects.
// A group of GraphicObjects is a GraphicObject itself and should be treated as such.
func TestComposite(t *testing.T) {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewSquare("Red"))
	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawing.Children = append(drawing.Children, group)

	// In this case, our drawing has several layers, but we are going to treat it as a single object using the String() method.
	s := drawing.String()
	assert.Contains(t, s, "My Drawing")
	assert.Contains(t, s, "*Yellow Circle")
	assert.Contains(t, s, "**Blue Square")
}
