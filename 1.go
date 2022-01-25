package main

import (
	"fmt"
	"math/rand"

	"time"
)

func getRandomNumber(a, b int) int{
	rand.Seed(time.Now().UnixNano())
	time.Sleep(10 * time.Millisecond)
	// a ≤ n ≤ b
	n := a + rand.Intn(b-a+1)
	return n
}

func matrixOutput(sample [4][4]int){
	for _, v := range sample {
		fmt.Println(v)
	}
}

func createRandomMatrix() [4][4]int{

	sample := [4][4]int{}


	for i, _ := range sample {
		for d, _ := range sample{
			sample[i][d] = getRandomNumber(0, 9)
		}
	}
	return sample
}

func main() {
	fmt.Println("Starting matrix:")
	sample := createRandomMatrix()
	matrixOutput(sample)
	
	n := len(sample)
	for i := 0; i < n / 2; i++ {
		for j := i; j < n - i - 1; j++ {

			y := sample[i][j]

			sample[i][j] = sample[n-1-j][i]

			sample[n - 1 - j][i] = sample[n - 1 - i][n - 1 - j]

			sample[n - 1 - i][n - 1 - j] = sample[j][n - 1 - i]

			sample[j][n - 1 - i] = y

		}
	}


	fmt.Print("\nResult:\n")
	matrixOutput(sample)

}
