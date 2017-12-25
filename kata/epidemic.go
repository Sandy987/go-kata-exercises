package kata

// Epidemic finds the largest number of infected in a hypothetical scenario posed in
// this kata: https://www.codewars.com/kata/disease-spread
func Epidemic(tm, n, s0, i0 int, b, a float64) int {
	dt := float64(tm) / float64(n)
	susceptible := float64(s0)
	infected := float64(i0)
	recovered := 0.0
	infectedMax := infected

	for k := 1; k < n; k++ {
		susceptibleNext := susceptible - (dt * b * susceptible * infected)
		infectedNext := infected + (dt * (b*susceptible*infected - (a * infected)))
		// While there is no strict need to track the recovered for this function, we do so for interest
		recoveredNext := recovered + (dt * infected * a)

		if infectedNext >= infectedMax {
			infectedMax = infectedNext
		}

		susceptible = susceptibleNext
		infected = infectedNext
		recovered = recoveredNext
	}

	return int(infectedMax)
}
