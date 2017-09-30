package kata

import "strconv"

var tape [30000]bool
var tp = 0

// Stack to keep track of branching instruction pointers
var branches []int

// Boolfuck is a function that interprets a boolfuck program according to the specifications here
// https://www.codewars.com/kata/esolang-interpreters-number-4-boolfuck-interpreter/train/go
func Boolfuck(code, input string) string {
	// Make definitively sure we're working with ASCII characters to enable predictable indexing
	code = strconv.QuoteToASCII(code)

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			tape[tp] = !tape[tp]
			break
		case '<':
			if tp == 0 {
				tp = 30000 - 1
			} else {
				tp = tp - 1
			}
			break
		case '>':
			if tp == 30000-1 {
				tp = 0
			} else {
				tp = tp + 1
			}
			break
		case ',':
			break
		case ';':
			break
		case '[':
			if tape[tp] == false {
				scan := i
				branchCounter := 1
				for branchCounter > 0 {
					scan++
					if code[scan] == '[' {
						branchCounter++
					} else if code[scan] == ']' {
						branchCounter--
					}
				}
				// Set instruction pointer to jump point
				i = scan
			} else {
				push(i)
			}
			break
		case ']':
			newPointer := pop()
			if newPointer != -1 {
				i = newPointer
			}
			break
		default:
			break
		}
	}

	return ""
}

// Functions to deal with the branching stack
func push(pointer int) {
	branches = append(branches, pointer)
}
func pop() int {
	if len(branches) == 0 {
		return -1
	}
	var p = branches[len(branches)-1]
	branches = branches[:len(branches)-1]
	return p
}
func peek() int {
	return branches[len(branches)-1]
}
