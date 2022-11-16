package main

import (
	"builder/reader"
	"fmt"
	"log"
)

func main() {
	c, err := reader.ReadConfig("C:/fork/core/silcker-config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", c)
}
