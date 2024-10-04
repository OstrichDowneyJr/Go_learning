package main

import (
	"fmt"
	c "strconv"
	s "strings"
)

var p = fmt.Println

func main() {
	var userInput string
	var exit int = 0
	p("hello")
	for exit <= 1 {
		p("what operation do you want to do?")
		fmt.Scan(&userInput)
		if s.Contains(userInput, "+") {
			var data = s.Split(userInput, "+")
			result1, err1 := c.Atoi(data[0])
			restul2, err2 := c.Atoi(data[1])
			if err1 != nil {
				panic(err1)
			}
			if err2 != nil {
				panic(err2)
			}

			p(result1 + restul2)
		}
		if s.Contains(userInput, "-") {
			var data = s.Split(userInput, "-")
			result1, err1 := c.Atoi(data[0])
			restul2, err2 := c.Atoi(data[1])

			if err1 != nil {
				panic(err1)
			}
			if err2 != nil {
				panic(err2)
			}

			p(result1 - restul2)
		}

		if s.Contains(userInput, "exit") {
			break
		}
	}
}
