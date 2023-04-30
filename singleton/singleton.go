package main

import (
	"sync"
)

type Database interface {
	GetValue(key string) int
}

type singletonDatabase struct {
	values map[string]int
}

func (db *singletonDatabase) GetValue(
	key string) int {
	return db.values[key]
}

var once sync.Once
var instance Database

func GetSingletonDatabase() Database {
	once.Do(func() {
		if instance == nil {
			instance = &singletonDatabase{}
		}
	})
	return instance
}
