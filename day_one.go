package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read_file(filename string) []byte {
  file, err := os.Open(filename)
  check(err)
  defer func() {
    err = file.Close()
    check(err)
  }()

  b, err := ioutil.ReadAll(file)
  check(err)
  return b
}

func main() {
  input := read_file("day_one_input.txt")

  fmt.Println(string(input))
}
