package fileOps

import (
	"fmt"
	"errors"
	"os"
	"strconv"
)

func ReadFloatFromFile(filename string) (float64,error) {
	data, err := os.ReadFile(filename)
	fmt.Println(data)
	if err != nil {
		return 1000,errors.New("Failed to read from the file!")
	}
	valueText := string(data)
	fmt.Println(valueText)
	value, err := strconv.ParseFloat(valueText,64)

	if err != nil {
		return 1000, errors.New("Failed to parse to float value!")
	}

	return value, nil
}

func WriteFloatToFile(filename string,value float64) error {
	valueText := fmt.Sprint(value)
	return os.WriteFile(filename, []byte(valueText), 0644)
}