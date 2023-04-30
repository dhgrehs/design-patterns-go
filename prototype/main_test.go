// The Prototype Pattern is a creational design pattern that allows an object to create a copy of itself,
//  which can then be modified without affecting the original object.
// The Prototype Pattern aims to clone an existing object instead of creating a new one.
// This is achieved by having a prototype interface that declares a method for cloning itself.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBadDeepCopy calls BadDeepCopy, update some values,
// and check valid return value.
func TestBadDeepCopy(t *testing.T) {
	john := BadCopiedPerson{Person{"John", &Address{
		"123 London Rd",
		"London",
		"UK"}, []string{}}}
	jane := john.DeepCopy()

	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"

	//This test will fail, because BadDeepCopy() is not a real deep copy.
	//It is a totally wrong implementation of a deep copy.
	//Jane is, in fact, a pointer to John.
	assert.NotEqual(t, john.Name, jane.Name, "The two names shouldn't be the same.")
	assert.NotEqual(t, john.Address.StreetAddress, jane.Address.StreetAddress, "The two addresses shouldn't be the same.")
}

// TestManualDeepCopy calls ManualDeepCopy, update some values,
// and check valid return value.
func TestManualDeepCopy(t *testing.T) {
	john := ManualCopiedPerson{Person{"John", &Address{
		"123 London Rd",
		"London",
		"UK"}, []string{"Chris"}}}
	jane := john.DeepCopy()

	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	//Even though we have a better implementation of a deep copy,
	//It is manually implemented, and it is not a good practice.
	assert.NotEqual(t, john.Name, jane.Name, "The two names shouldn't be the same.")
	assert.NotEqual(t, john.Address.StreetAddress, jane.Address.StreetAddress, "The two addresses shouldn't be the same.")
	assert.NotEqual(t, john.Friends, jane.Friends, "The two friend lists shouldn't be the same.")
}

// TestMDeepCopy calls DeepCopy, update some values,
// and check valid return value.
func TestDeepCopy(t *testing.T) {
	john := Person{"John", &Address{
		"123 London Rd",
		"London",
		"UK"}, []string{"Chris"}}
	jane := john.DeepCopy()

	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	//The DeepCopy implementation used here is a generic implementation using binary encoding.
	//If a new field is added to the Person struct, the DeepCopy implementation will still work.
	assert.NotEqual(t, john.Name, jane.Name, "The two names shouldn't be the same.")
	assert.NotEqual(t, john.Address.StreetAddress, jane.Address.StreetAddress, "The two addresses shouldn't be the same.")
	assert.NotEqual(t, john.Friends, jane.Friends, "The two friend lists shouldn't be the same.")
}
