package main

import (
  "fmt"
  fb "./fizzBuzz"
)

func main() {
  arr := fb.FizzBuzz(0, 15)
  fmt.Printf("%v", arr)
}
