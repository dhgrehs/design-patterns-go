// Singleton is a creational design pattern,
// which ensures that only one object of its kind exists and provides a single point of access to it for any other code.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSimpleSingletonDatabase calls GetSimpleSingletonDatabase and gets a value.
func TestGetSimpleSingletonDatabase(t *testing.T) {
	db := GetSimpleSingletonDatabase()

	// This simple singleton works fine, however it breaks a SOLID principle.
	// You depend directly on the concrete implementation of the database.
	// It is not a good practice and the Dependency Inversion Principle is broken.
	v := db.GetValue("something")
	assert.Equal(t, 123, v, "The value should be 123.")
}

func TestGetSingletonDatabase(t *testing.T) {
	db := GetSingletonDatabase()

	// This simple singleton also works fine, but instead of depending on the concrete implementation of the database,
	// it depends on an interface. This is a better implementation, because it follows the Dependency Inversion Principle.

	// It means that we can replace the concrete implementation of the database with another one, and the code will still work.
	// For example, a dummy database for testing purposes.
	v := db.GetValue("anything")
	assert.Equal(t, 0, v, "The value should be 0.")
}
