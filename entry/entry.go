package main

import (
  "fmt"
  "krzysztofsan/krzygen"
)

func main() {
  message := krzygen.Generate()

  fmt.Println(message)
}
