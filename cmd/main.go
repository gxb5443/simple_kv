package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gxb5443/pathgather-kv/store"
)

func main() {

	store := new(store.StackStore)
	store.Initialize()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "READ":
			if len(line) != 2 {
				fmt.Println("Incorrect Usage")
				fmt.Println("USAGE: READ <key>")
				continue
			}
			fmt.Println("READING KEY " + line[1])
			val, err := store.Read(line[1])
			fmt.Println(val)
			if err != nil {
				fmt.Println(err)
			}
		case "WRITE":
			if len(line) != 3 {
				fmt.Println("Incorrect Usage")
				fmt.Println("USAGE: WRITE <key> <value>")
				continue
			}
			fmt.Println("WRITING VALUE " + line[2] + " To KEY " + line[1])
			store.Write(line[1], line[2])
		case "DELETE":
			if len(line) != 2 {
				fmt.Println("Incorrect Usage")
				fmt.Println("USAGE: DELETE <key>")
				continue
			}
			fmt.Println("DELETING KEY " + line[1])
		case "START":
			fmt.Println("STARTING Transaction" + line[0])
		case "COMMIT":
			fmt.Println("COMMITING Transaction" + line[0])
		case "ABORT":
			fmt.Println("ABORTING Transaction" + line[0])
		case "QUIT":
			return
		default:
			fmt.Println("Unknown Command " + line[0])
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
