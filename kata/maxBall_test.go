package kata

import "testing"

func TestMaxBall(t *testing.T) {
	var result int
	result = MaxBall(37)
	if result != 10 {
		t.Error("Failed Test for ", 37, "Got ", result)
	}

	result = MaxBall(45)
	if result != 13 {
		t.Error("Failed Test for ", 45, "Got ", result)
	}

	result = MaxBall(99)
	if result != 28 {
		t.Error("Failed Test for ", 99, "Got ", result)
	}
	result = MaxBall(85)
	if result != 24 {
		t.Error("Failed Test for ", 85, "Got ", result)
	}
	result = MaxBall(136)
	if result != 39 {
		t.Error("Failed Test for ", 136, "Got ", result)
	}
}
