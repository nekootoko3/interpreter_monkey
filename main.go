package main

import (
	"fmt"
	"monkey/repl"
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
	fmt.Printf("やってきなよyou!\n")
	repl.Start(os.Stdin, os.Stdout)
}
