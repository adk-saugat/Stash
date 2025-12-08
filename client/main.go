package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adk-saugat/stash/commands"
)

func main() {
	if len(os.Args) < 2{
		log.Fatal("no commands found")
	}
	
	switch (os.Args[1]){
	case "login":
		commands.Login()
	case "register":
		commands.Register()
	case "create":
		commands.Create()
	case "help":
		fmt.Println("help")
	default:
		log.Fatal("no command found... use help for command info")
	}
}