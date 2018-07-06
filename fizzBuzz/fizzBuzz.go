package fizzBuzz

import (
  "strconv"
)

func FizzBuzz(start int, finish int) []string {
  var arr []string

  for i := start; i <= finish; i++ {
    arr = append( arr, chechNumber(i) )
  }

  return arr
}

func chechNumber(num int) string {
  f := "fizz"
  b := "buzz"

  if (num == 0) {
    return "0"
  }

  if (num%15 == 0) { 
    return f+b
  }
  if (num%3 == 0) {
    return f
  }
  if (num%5 == 0) {
    return b
  }

  return strconv.Itoa(num)
}
