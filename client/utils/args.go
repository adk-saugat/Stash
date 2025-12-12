package utils

import (
	"log"
	"os"
)

func GetArg(index int, errorMsg string) string {
	actualIndex := index + 1
	if len(os.Args) < actualIndex+1 || os.Args[actualIndex] == "" {
		log.Fatal(errorMsg)
	}
	return os.Args[actualIndex]
}

