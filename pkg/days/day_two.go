package days

import (
	"fmt"
	"strings"

	"github.com/laukei/advent_of_code_2022/pkg/utils"
)

type Choice int

const (
  Rock = iota + 1
  Paper
  Scissors
)
const Loss = 0
const Draw = 3
const Win = 6

// Determines the outcome of the game
func determineOutcome(input_move Choice, response_move Choice) int {
  if input_move == response_move { 
    return Draw
  }

  // If they're not the same, then it's either a win or a loss
  if ((response_move - input_move) % 3) == 1 {
    return Win
  } else {
    return Loss
  }
}

func DayTwo(fn string) {
  input_map := map[string]Choice{
    "A": Rock,
    "B": Paper,
    "C": Scissors,
  }
  response_map := map[string]Choice{
    "X": Rock,
    "Y": Paper,
    "Z": Scissors,
  }

  input := string(utils.ReadFile(fn))
  s := strings.Split(input, "\n")
  fmt.Printf(s[0])
  
  for _, r := range s {
    moves := strings.Split(r, " ")
    input_move := input_map[moves[0]]
    response_move := response_map[moves[1]]
    fmt.Printf("input move %s and response move %s", moves[0], moves[1])
    fmt.Println(determineOutcome(input_move, response_move))
  }
}
