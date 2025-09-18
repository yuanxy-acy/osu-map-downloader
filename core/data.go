package core

import (
	"fmt"
	"io"
	"os"
)

var dataPath = "data/"

func saveData(fileName string, data *[]byte) {
	file, err := os.OpenFile(dataPath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)
	_, err = file.Write(*data)
	if err != nil {
		fmt.Println(err)
	}
}

func readData(fileName string) []byte {
	file, err := os.OpenFile(dataPath+fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
