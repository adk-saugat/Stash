package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2{
		log.Fatal("no commands found")
	}
	
	switch (os.Args[1]){
	case "login":
		fmt.Println("login")
	case "register":
		fmt.Println("register")
	case "help":
		fmt.Println("help")
	default:
		log.Fatal("no command found... use help for command info")
	}
}