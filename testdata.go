package main

import "math/rand"

func generateTestData(fields, size int) []interface{} {
	id := rand.Int()

	values := make([]interface{}, fields)

	values = append(values, id)

	return values
}
