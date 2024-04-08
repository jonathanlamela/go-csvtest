//main2.go

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

type Record struct {
	Brand          string  `csv:"ManufacturerName"`
	MPN            string  `csv:"Number"`
	EAN13          string  `csv:"EAN"`
	Name           string  `csv:"ArticleFullName"`
	Description    string  `csv:"Description"`
	WholesalePrice float32 `csv:"PricePreTax"`
	IsAvailable    bool    `csv:"Stock"`
}

func main() {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		r.LazyQuotes = true
		return r // Allows use pipe as delimiter
	})

	// Open the CSV file
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var records []Record
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	// Print the records
	for _, record := range records {
		fmt.Println("Name", "\t\t", record.Name)
		fmt.Println("Brand", "\t\t", record.Brand)
		fmt.Println("Price", "\t\t", record.WholesalePrice)
		fmt.Println("Available", "\t", record.IsAvailable)
		fmt.Println()
	}
}
