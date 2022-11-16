package main

import (
   "fmt"
   "github.com/silcker/builder/reader/reader"
)

func main() {
   fmt.Printf(reader.Config("C:/fork/core/silcker-config.yaml") );
}