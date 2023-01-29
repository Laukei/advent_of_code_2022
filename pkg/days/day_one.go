package days

import (
	"fmt"
	"strings"
  "github.com/laukei/advent_of_code_2022/pkg/utils"
)

func DayOne() {
  input := string(utils.ReadFile("day_one_input.txt"))
  s := strings.Split(input, "\n")
  totals := utils.Totalize(s)
  max := utils.Max(totals)

  fmt.Printf("The elf carrying the largest number of calories is carrying %d calories.", max)

  utils.Sort(totals)
  top_3 := utils.SumFirstN(totals, 3)
  fmt.Printf("The top three elves are jointly carrying %d calories.", top_3)
}
