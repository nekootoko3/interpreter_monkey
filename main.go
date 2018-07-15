package main

import (
	"fmt"
	"interpreter_monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the SAS LOVE Programming language!\n",
		user.Username)
	fmt.Printf("Ginobili! Duncan! Parker!\n")
	repl.Start(os.Stdin, os.Stdout)
}
