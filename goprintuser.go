package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {

	uid := syscall.Geteuid()
	gid := syscall.Getegid()

	usr, err := user.LookupId(strconv.Itoa(uid))

	if err != nil {
		fmt.Println("Error: Not able to get current user.")
		os.Exit(1)
	}

	fmt.Printf("current user: euid=%d egid=%d username=%s\n",
		uid,
		gid,
		usr.Username,
	)

}
