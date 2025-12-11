package commands

import (
	"log"
	"os"
)

func Create(){
	// creating the .stash folder
	exists, err := folderExists(".stash")
	if err != nil {
		log.Fatal("Error: Couldnot check folders existence.")
	}

	if !exists {
		err = os.Mkdir(".stash", 0755)
		if err != nil {
			log.Fatal("Error: Couldnot create folder.")
		}
	}

	// creating the store folder
	exists, err = folderExists(".stash/store")
	if err != nil {
		log.Fatal("Error: Couldnot check folders existence.")
	}

	if !exists {
		err = os.Mkdir("./.stash/store", 0755)
		if err != nil {
			log.Fatal("Error: Couldnot create.")
		}
	}
}

func folderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil 
	}
	if os.IsNotExist(err) {
		return false, nil 
	}
	return false, err 
}
