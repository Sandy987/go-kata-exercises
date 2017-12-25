package kata

import "fmt"

var annMap = make(map[int]int)
var johnMap = make(map[int]int)

func Ann(n int) []int {
	fmt.Printf("Ann %d\n", n)
	katas := []int{1}
	if n == 0 {
		annMap[0] = 1
		return katas
	}
	annMap[1] = 1
	katas = append(katas, 1)
	for i := 1; i < n; i++ {
		doneYtdy, ok := annMap[i-1]
		if !ok {
			anns := Ann(i - 1)
			doneYtdy = anns[len(anns)-1]
			annMap[i-1] = doneYtdy
			fmt.Printf("Wrote %d into Ann[%d]\n", doneYtdy, i-1)
		}

		johnsDone, ok := johnMap[doneYtdy]
		if !ok {
			johns := John(doneYtdy)
			johnsDone = johns[len(johns)-1]
			johnMap[doneYtdy] = johnsDone
			fmt.Printf("Wrote %d into John[%d]\n", johnsDone, doneYtdy)
		}

		doneToday := i - johnsDone

		katas = append(katas, doneToday)
	}
	return katas
}
func John(n int) []int {
	fmt.Printf("John %d\n", n)
	katas := []int{0}
	if n == 0 {
		johnMap[0] = 0
		return katas
	}
	johnMap[1] = 0
	katas = append(katas, 0)
	for i := 1; i < n; i++ {
		doneYtdy, ok := johnMap[i-1]
		if !ok {
			johns := John(i - 1)
			doneYtdy = johns[len(johns)-1]
			johnMap[i-1] = doneYtdy
			fmt.Printf("Wrote %d into John[%d]\n", doneYtdy, i-1)
		}

		annsDone, ok := annMap[doneYtdy]
		if !ok {
			anns := Ann(doneYtdy)
			annsDone = anns[len(anns)-1]
			annMap[doneYtdy] = annsDone
			fmt.Printf("Wrote %d into Ann[%d]\n", annsDone, doneYtdy)
		}

		doneToday := i - annsDone

		katas = append(katas, doneToday)
	}
	return katas
}
func SumJohn(n int) int {
	johns := John(n)
	sum := 0
	for i := range johns {
		sum += i
	}
	return sum
}
func SumAnn(n int) int {
	anns := Ann(n)
	sum := 0
	for i := range anns {
		sum += i
	}
	return sum
}
