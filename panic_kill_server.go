package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't open pid file (is the Server running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("Error: the file cannot be removed. %s", err)
	}

	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid)

	if err != nil {
		return errors.Wrap(err, "Bad process Id")
	}

	fmt.Printf("Killing process id %d\n", pid)

	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(10)
	}
}
