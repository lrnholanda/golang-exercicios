package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// Define the data
	records := [][]string{
		{"John", "Doe", "35"},
		{"Jane", "Doe", "29"},
		{"Jim", "Smith", "41"},
	}
	filename := "users.csv"

	// Write it to CSV
	err := WriteToCsv(records, filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created file:", filename)
}

// WriteToCsv accepts a slice of strings to write to a csv file.
func WriteToCsv(records [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	return nil
}
