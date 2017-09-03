package kata

import (
	"bytes"
)

// Interpreter takes in a MiniStringFuck program as string and outputs its results.
func Interpreter(code string) string {
	var memcell byte
	var output bytes.Buffer
	for _, c := range code {
		if c == '+' {
			memcell = memcell + 1
		} else if c == '.' {
			output.WriteByte(memcell)
		}
	}
	return output.String()
}
