package kata

//Solution for https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go

var vals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Decode take a string containing roman numerals and converts it to a decimal number.
func Decode(roman string) int {
	var sum = 0
	for i := 0; i < len(roman); i++ {
		var current = string(roman[i])
		switch current {
		case "M", "D", "L", "V":
			sum += vals[current]
		case "I", "X", "C":
			var adder = vals[current]
			if i+1 < len(roman) {
				var next = string(roman[i+1])
				if isSubtractiveNotation(current, next) {
					adder = vals[next] - vals[current]
					i = i + 1 //skip the next numeral
				}
			}
			sum += adder
		}
	}

	return sum
}

func isSubtractiveNotation(current, next string) bool {
	switch current {
	case "I":
		return next == "V" || next == "X"
	case "X":
		return next == "L" || next == "C"
	case "C":
		return next == "D" || next == "M"
	default:
		return false
	}
}
