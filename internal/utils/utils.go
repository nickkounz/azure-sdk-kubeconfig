package utils

import (
	"fmt"
	"os"
)

func HomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return dirname
}
