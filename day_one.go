package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Error handling
func check(e error) {
  if e != nil {
    panic(e)
  }
}

// Reads filename fn and returns the file's bytes
func ReadFile(fn string) []byte {
  file, err := os.Open(fn)
  check(err)
  defer func() {
    err = file.Close()
    check(err)
  }()

  b, err := ioutil.ReadAll(file)
  check(err)
  return b 
}

// Totalizes the contiguous numbers
func Totalize(input []string) []int {
  var result []int
  counter := 0
  for _, row := range input {
    if row == "" {
      result = append(result, counter)
      counter = 0
    } else {
      value, err := strconv.Atoi(row)
      check(err)
      counter += value
    }
  }

  return result
}

// Get maximum value
func Max(input []int) int {
  result := 0

  for _, row := range input {
    if row > result {
      result = row
    }
  }

  return result
}

// Sort highest to lowest
func Sort(input []int) {
  sort.Slice(input, func(i, j int) bool {
    return input[i] > input[j]
  })
}

// Sums the first N elements
func SumFirstN(input []int, n int) int {
  result := 0

  for i := 0; i < n; i++ {
    result += input[i]
  }

  return result
}

func main() {
  input := string(ReadFile("day_one_input.txt"))
  s := strings.Split(input, "\n")
  totals := Totalize(s)
  max := Max(totals)

  fmt.Printf("The elf carrying the largest number of calories is carrying %d calories.", max)

  Sort(totals)
  top_3 := SumFirstN(totals, 3)
  fmt.Printf("The top three elves are jointly carrying %d calories.", top_3)
}
