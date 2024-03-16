package main

import (
	"encoding/csv"
	"github.com/google/uuid"
	"os"
)

func main() {
	// Create a file to write CSV data
	file, err := os.Create("uuids.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	err = writer.Write([]string{"UUID"})
	if err != nil {
		panic(err)
	}

	// Generate and write 100,000 unique UUIDs
	for i := 0; i < 50000; i++ {
		newUUID := uuid.New().String()
		err := writer.Write([]string{newUUID})
		if err != nil {
			panic(err)
		}
	}
}
