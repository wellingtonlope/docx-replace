package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/nguyenthenguyen/docx"
	"github.com/urfave/cli/v2"
)

func main() {
	var (
		templateFileName string
		dataFileName     string
		separator        string
	)

	app := &cli.App{
		Name:  "docx-replace",
		Usage: "A simple command-line project to use a .csv database to create some .docx documents based on a template.A simple project in command line to use a database in csv to create.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "template",
				Aliases:     []string{"t"},
				Usage:       "to specify the .docx template",
				Required:    true,
				Destination: &templateFileName,
			},
			&cli.StringFlag{
				Name:        "data",
				Aliases:     []string{"d"},
				Usage:       "to specify the .csv data",
				Required:    true,
				Destination: &dataFileName,
			},
			&cli.StringFlag{
				Name:        "sepator",
				Value:       ",",
				Aliases:     []string{"s"},
				Usage:       "to specify the separator in .csv data",
				Required:    false,
				Destination: &separator,
			},
		},
		Action: func(c *cli.Context) error {
			replaceableDoc, err := docx.ReadDocxFile(templateFileName)
			if err != nil {
				return fmt.Errorf("unable to read template file %s", templateFileName)
			}
			defer replaceableDoc.Close()

			csvFile, err := readCsvFile(dataFileName, separator)
			if err != nil {
				return err
			}
			headers := csvFile[0]

			for _, line := range csvFile[1:] {
				doc := replaceableDoc.Editable()
				for indice, column := range line {
					doc.Replace(headers[indice], column, -1)
				}
				outPutDocx := fmt.Sprintf("%s.docx", line[0])
				doc.WriteToFile(outPutDocx)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func readCsvFile(filePath, separator string) ([][]string, error) {
	if len(separator) != 1 {
		return nil, fmt.Errorf("the separator %q is invalid", separator)
	}
	separatorRune := []rune(separator)[0]

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s", filePath)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = separatorRune
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse file as CSV for %s", filePath)
	}

	return records, nil
}
