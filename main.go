package main

import (
   "fmt"
   "github.com/silcker/builder/reader"
)

func main() {
   fmt.Printf(reader.Config("C:/fork/core/silcker-config.yaml") );
}