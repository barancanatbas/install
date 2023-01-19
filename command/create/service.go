package create

import (
	"errors"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(fileName string, data []byte) error {
	if data == nil {
		return errors.New("data is null")
	}

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}
