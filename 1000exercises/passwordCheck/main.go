package main

import (
	"fmt"
	"os"
)

const(
	user1 = "jack"
	user2 = "inanc"

	pass1 = "1888"
	pass2 = "1879"

	usage = "Usage: [username] [password]"
	errUser = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
)


func main () {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u == user1 && p == pass1 || u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else if u == user1 && p != pass1 || u == user2 && p != pass2 {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(errUser, u)
	}
}