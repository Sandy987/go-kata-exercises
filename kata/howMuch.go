package kata

import "strconv"

// HowMuch attempts to find a value between m, n that satisfies luxury purchasing rules
func HowMuch(m, n int) [][3]string {
	var low, high int
	var possibles [][3]string

	// Get the low and high bounds of how much money is possible
	if m <= n {
		low = m
		high = n
	} else {
		low = n
		high = m
	}

	// For each possible value of money
	for i := low; i <= high; i++ {
		// Check if it's possible to satisfy both sets of purchases
		if ((i-1)%9 == 0) && ((i-2)%7 == 0) {
			costCar := (i - 1) / 9
			costBoat := (i - 2) / 7
			possibility := [3]string{
				"M: " + strconv.Itoa(i),
				"B: " + strconv.Itoa(costBoat),
				"C: " + strconv.Itoa(costCar),
			}
			possibles = append(possibles, possibility)
		}
	}
	return possibles
}
