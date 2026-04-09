package helper

import (
	"os"

	"github.com/gocarina/gocsv"
)

func SaveCSVFile(pathfile string, data any) error {
	file, err := os.OpenFile(pathfile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(data, file)
}
