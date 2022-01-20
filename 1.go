package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(a, b int) int{
	rand.Seed(time.Now().UnixNano())
	time.Sleep(10 * time.Millisecond)
	n := a + rand.Intn(b-a+1) // a ≤ n ≤ b
	return n
}

func matrixOutput(sample [5][5]int){
	for _, v := range sample {
		fmt.Println(v)
	}
}

func createRandomMatrix() [5][5]int{
	/*

	Идея заключалась в том, чтобы размерность массива вводил пользователь,
	но с двумерным массивом возникли проблемы

	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter numbers of columns: ")

	sc.Scan()
	p := sc.Text()

	fmt.Println("Enter numbers of rows: ")
	sc.Scan()
	q := sc.Text()

	colNum, _ := strconv.Atoi(p)
	rowNum, _ := strconv.Atoi(q)

	sample := make([][]int, colNum, rowNum)
	*/



	sample := [5][5]int{}

	for i, _ := range sample {//заполнение рандомными числами, 0-9 для красоты матрицы при выводе
		for d, _ := range sample{
			sample[i][d] = getRandomNumber(0, 9)
		}
	}
	return sample
}

func main() {
	fmt.Println("Starting matrix:\n")
	sample := createRandomMatrix()
	matrixOutput(sample)


	newSample := [5][5]int{}
	for i, _ := range newSample {	//для запоминания нулей основной матрицы
		for d, _ := range newSample{
			newSample[i][d] = 1
		}
	}

	for i := 0; i < cap(sample); i++ {	//запоминаем нули во вторую матрицу
	//for i, v := range sample{
		for d := 0; d < cap(sample); d++{
			if sample[i][d] == 0{
				newSample[i][d] = sample[i][d]
			}
		}
	}

	for i, _ := range newSample{	//ищем нуль во вспомогательной
		for d := 0; d < cap(newSample); d++{
			if newSample[i][d] == 0 {	//если нашли

				for z, _ := range sample {	//все элементы при совпадении строки или столбца обращаем в нуль
					for x := 0; x < cap(newSample); x++ {

						if i == z {
							sample[z][x] = 0

						} else if d == x {
							sample[z][x] = 0
						}
					}
				}
			}

		}
	}

	fmt.Print("\nSaved zeros:\n")
	matrixOutput(newSample)
	fmt.Print("\nResult:\n")
	matrixOutput(sample)

}
