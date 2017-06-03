package kata

import "testing"

func TestLongestConcec(t *testing.T) {
	var result = ""
	result = LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2)
	if result != "abigailtheta" {
		t.Error("Failed Test for ", "abigailtheta", "Got ", result)
	}

	result = LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1)
	if result != "oocccffuucccjjjkkkjyyyeehh" {
		t.Error("Failed Test for ", "oocccffuucccjjjkkkjyyyeehh", "Got ", result)
	}

	result = LongestConsec([]string{}, 3)
	if result != "" {
		t.Error("Failed Test for empty", "Got ", result)
	}
	result = LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2)
	if result != "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck" {
		t.Error("Failed Test for ", "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck", "Got ", result)
	}
}
