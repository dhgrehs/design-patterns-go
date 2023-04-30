// Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.
// In this example, we have a VectorImage interface that represents a vector image.
// We also have a RasterImage interface that represents a raster (pixelated) image.
// We want to be able to draw raster images on the screen, but we only have vector images.
// So we create an adapter that allows us to draw vector images on the screen.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem to solve: I want to print a RasterImage, but I can only make a VectorImage
// Solution: Create an adapter that adapts a VectorImage to a RasterImage
func TestAdapter(t *testing.T) {
	rect := NewRectangle(6, 4)
	rasterRect := VectorToRaster(rect) // adapter!

	// Validates that we are, in fact, adapting a vector image to a raster image
	_, isRasterImage := rasterRect.(RasterImage)
	assert.True(t, isRasterImage)

	// Validates that the raster image has the correct number of points
	assert.NotEmpty(t, rasterRect.GetPoints())
}
