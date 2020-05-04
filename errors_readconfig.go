package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
}

func readConfig(path string) (*Config, error) {
	_, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error in readine configuration file")
	}

	config := &Config{}

	return config, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("/ope/test.cfg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		log.Printf("Error: %+v", err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
