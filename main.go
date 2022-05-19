/* TODO(tijani):
 
 - add quit command to the repl.
 - remove semicolons from the language to mark the end of a statement.
 - add line and position reporting in the lexer when a lexing or parsing error occurs.
*/


package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}