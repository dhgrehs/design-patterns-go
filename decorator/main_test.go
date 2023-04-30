// Decorator pattern is a structural design pattern that lets you attach new behaviors to objects.
// The decorator pattern is an alternative to subclassing.
// And also, the Open Closed Principle is respected.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleAggregation(t *testing.T) {

	d := MultipleAggregationDragon{}
	// You cannot use d.Age because it is ambiguous
	//d.Age = 10
	// You can use Setter and Getter, but you still may face issues if the age is set directly
	d.SetAge(10)

	assert.Contains(t, d.Fly(), "Simple Bird Flying!")

	// For example, how do you know that the age is set correctly?
	// Flying and crawling looks wrong.
	d.SimpleLizard.Age = 8
	assert.Contains(t, d.Crawl(), "Simple Lizard Crawling!")

}

func TestDecorator(t *testing.T) {
	d := NewDragon()

	d.SetAge(10)
	assert.Contains(t, d.Fly(), "Bird Flying!")
	assert.NotContains(t, d.Crawl(), "Lizard Crawling!")
}
