package test

import (
  "testing"
  "reflect"

  fb "../fizzBuzz"
)

const _f  string = "fizz"
const _b  string = "buzz"
const _fb string = _f + _b

func TestZero(t *testing.T) {
  mustBe := []string{ "0" }
  output := fb.FizzBuzz(0, 0)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestNum43(t *testing.T) {
  mustBe := []string{ "43" }
  output := fb.FizzBuzz(43, 43)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestNum75(t *testing.T) {
  mustBe := []string{ _fb }
  output := fb.FizzBuzz(75, 75)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestNum100(t *testing.T) {
  mustBe := []string{ _b }
  output := fb.FizzBuzz(100, 100)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestFrom1to5(t *testing.T) {
  mustBe := []string{ "1", "2", _f, "4", _b }
  output := fb.FizzBuzz(1, 5)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}


func TestFrom14to21(t *testing.T) {
  mustBe := []string{ "14", _fb, "16", "17", _f, "19", _b, _f }
  output := fb.FizzBuzz(14, 21)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestFrom55to60(t *testing.T) {
  mustBe := []string{ _b, "56", _f, "58", "59", _fb }
  output := fb.FizzBuzz(55, 60)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestFrom95to99(t *testing.T) {
  mustBe := []string{ _b, _f, "97", "98", _f }
  output := fb.FizzBuzz(95, 99)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, got %v", mustBe, output)
  }
}

func TestFinal1to100(t *testing.T) {
  mustBe := []string {  
    "1",  "2",  _f, "4",  _b, _f, "7",  "8",  _f, _b, "11", _f, "13", "14", _fb,
    "16", "17", _f, "19", _b, _f, "22", "23", _f, _b, "26", _f, "28", "29", _fb,
    "31", "32", _f, "34", _b, _f, "37", "38", _f, _b, "41", _f, "43", "44", _fb,
    "46", "47", _f, "49", _b, _f, "52", "53", _f, _b, "56", _f, "58", "59", _fb,
    "61", "62", _f, "64", _b, _f, "67", "68", _f, _b, "71", _f, "73", "74", _fb,
    "76", "77", _f, "79", _b, _f, "82", "83", _f, _b, "86", _f, "88", "89", _fb,
    "91", "92", _f, "94", _b, _f, "97", "98", _f, _b }
  output := fb.FizzBuzz(1, 100)

  if( !reflect.DeepEqual(mustBe, output) ) {
    t.Errorf("Expected %v, \ngot %v", mustBe, output)
  }
}

