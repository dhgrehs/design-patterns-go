package main

import "sync"

type simpleSingletonDatabase struct {
	values map[string]int
}

func (db *simpleSingletonDatabase) GetValue(key string) int {
	return db.values[key]
}

var simpleSingletonOnce sync.Once
var simpleSingletonDatabaseInstance *simpleSingletonDatabase

func GetSimpleSingletonDatabase() *simpleSingletonDatabase {
	simpleSingletonOnce.Do(func() {
		if simpleSingletonDatabaseInstance == nil {
			simpleSingletonDatabaseInstance = &simpleSingletonDatabase{}
			// Initialize the map with sample data for testing.
			simpleSingletonDatabaseInstance.values = make(map[string]int)
			simpleSingletonDatabaseInstance.values["something"] = 123
		}
	})
	return simpleSingletonDatabaseInstance
}
