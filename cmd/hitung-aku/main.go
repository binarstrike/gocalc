package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/binarstrike/gocalc/v2"
)

var (
	operatorPatterRegex = regexp.MustCompile("^[*/+-]$")
	flagPatternRegex    = regexp.MustCompile("^-{1,2}[a-zA-Z]+$")
)

func main() {
	args := os.Args[1:]
	var interactiveMode bool

	if a := len(args); a == 0 {
		interactiveMode = true
	} else if flagPatternRegex.Match([]byte(args[0])) {
		flag.BoolVar(&interactiveMode, "i", false, "Enter intercative mode")
		flag.Parse()
	} else if a%2 == 0 {
		fmt.Print("Usage: hitung-aku 5 * 5 + 10\nto enter interactive mode pass flag -i or just run command with empty parameter")
		return
	}

	if interactiveMode {
		fmt.Println("Simple Calculator\ntype q for quit from application\nexample input: 5 * 2 / 2 + 3 + 15")
		enterInteractivePromptMode()
		return
	}

	operators, nums, err := parseInputArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := calculateResult(nums, operators)
	if err != nil {
		fmt.Printf("last result: %v\nerror: %v\n", result, err)
		return
	}

	fmt.Println(result)
}

func parseInputArgs(args []string) (operators []string, nums []float64, err error) {
	for i := 0; i < len(args); i += 2 {
		num, parseError := strconv.ParseFloat(args[i], 64)
		if parseError != nil {
			err = fmt.Errorf("failed to parsing number, invalid parameter number \"%s\"", args[i])
			break
		}
		nums = append(nums, num)

		if i == len(args)-1 {
			break
		}

		if opr := args[i+1]; operatorPatterRegex.Match([]byte(opr)) {
			operators = append(operators, opr)
		} else {
			err = fmt.Errorf("invalid parameter operator \"%s\"", opr)
			break
		}
	}
	return
}

func calculateResult[T gocalc.Number](nums []T, operators []string) (T, error) {
	c := gocalc.New(nums[0])

	for i, opr := range operators {
		switch opr {
		case "+":
			c.Plus(nums[i+1])
		case "-":
			c.Minus(nums[i+1])
		case "*":
			c.Multiply(nums[i+1])
		case "/":
			c.Divide(nums[i+1])
		}
	}

	return c.ResultAndError()
}

func enterInteractivePromptMode() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		s.Scan()

		input := strings.Split(s.Text(), " ")

		if input[0] == "q" {
			fmt.Println("Bye")
			return
		} else if a := len(input); a <= 1 {
			fmt.Println(input[0])
			continue
		} else if a%2 == 0 {
			fmt.Println("invalid input parameter")
			continue
		}

		operators, nums, err := parseInputArgs(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculateResult(nums, operators)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}
