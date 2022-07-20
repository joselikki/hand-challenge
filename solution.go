package main

import (
	"fmt"
)

func removeByIndex(r []int, index int) []int {
	return append(r[:index], r[index+1:]...)
}

func increment(n *int) {
	if *n == 255 {
		*n = 0
	} else {
		*n += 1
	}
}

func decrease(n *int) {
	if *n == 0 {
		*n = 255
	} else {
		*n -= 1
	}
}

func main() {
	input := "👉👆👆👆👆👆👆👆👆🤜👇👈👆👆👆👆👆👆👆👆👆👉🤛👈👊👉👉👆👉👇🤜👆🤛👆👆👉👆👆👉👆👆👆🤜👉🤜👇👉👆👆👆👈👈👆👆👆👉🤛👈👈🤛👉👇👇👇👇👇👊👉👇👉👆👆👆👊👊👆👆👆👊👉👇👊👈👈👆🤜👉🤜👆👉👆🤛👉👉🤛👈👇👇👇👇👇👇👇👇👇👇👇👇👇👇👊👉👉👊👆👆👆👊👇👇👇👇👇👇👊👇👇👇👇👇👇👇👇👊👉👆👊👉👆👊"

	runeInput := []rune(input)
	data := make([]int, 10)
	loops := make(map[int]int)
	stack := []int{}

	for i, inst := range runeInput {
		if string(inst) == "🤜" {
			stack = append(stack, i)
		} else if string(inst) == "🤛" {
			startLoop := stack[len(stack)-1]
			loops[startLoop] = i
			loops[i] = startLoop
			stack = removeByIndex(stack, len(stack)-1)
		}
	}

	pointer := 0
	cursor := 0

	for cursor < len(runeInput) {
		curInst := string(runeInput[cursor])

		switch curInst {
		case "👉":
			pointer += 1
		case "👈":
			pointer -= 1
		case "👆":
			increment(&data[pointer])
		case "👇":
			decrease(&data[pointer])
		case "🤜":
			if data[pointer] == 0 {
				cursor = loops[cursor]
			}
		case "🤛":
			if data[pointer] != 0 {
				cursor = loops[cursor]
			}
		case "👊":
			asciiCode := data[pointer]
			fmt.Printf("%s", string(rune(asciiCode)))
		}
		cursor += 1
	}

}
