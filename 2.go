package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var newLetters []string
	var newLetters string
	var k int = 1

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	word := sc.Text()

	letters := strings.Split(word, "")
	letters = append(letters, "")
	fmt.Println(letters)

	for i := 0; i < len(letters); i++ {
		if letters[i] == ""{
			break
		} else if letters[i] == letters[i+1] {
			k++
		}else {
			//fmt.Print(letters[i], k)
			s := string(k)
			newLetters += letters[i] + s
			k=1
		}

	}

	if len(newLetters) < len(letters) {
		fmt.Println(newLetters)
	} else {
		for _, v := range letters{
			fmt.Print(v)
		}
	}
	
/*почему-то выводит спец. символы в консоль вместо чисел в L38.
чтобы понять правильный ли вывод, закомменитить L37-43, разкомментить L29
 */

}