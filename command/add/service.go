package add

import (
	"encoding/json"
	"fmt"
	"github.com/barancanatbas/install/model"
	"os"
)

func readFile(fileName string) ([]model.Dependency, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var response []model.Dependency
	err = json.Unmarshal(data, &response)
	return response, err
}

func writeFile(fileName string, data []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, "%s\n", data); err != nil {
		fmt.Println(err)
	}
	return nil
}
