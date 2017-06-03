package kata

// MaxBall returns the highest height of a ball thrown from an initial velocity v0, recorded by a machine that records in 1/10 second increments
func MaxBall(v0 int) int {
	vms := float64(v0) / 3.6 // Get v0 in m/s
	curHeight := 0.0
	for i := 1; ; i++ {
		t := float64(i) / 10.0
		newHeight := (vms * t) - (0.5 * 9.81 * t * t)
		if newHeight < curHeight { // If we dropped height, return the previous record
			return i - 1
		}
		curHeight = newHeight
	}
}
