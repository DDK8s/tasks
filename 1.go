package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var notUnique bool
	fmt.Printf("Enter a word: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	word := sc.Text()

	letters := strings.Split(word, "")
	fmt.Println(letters)

	for i, v := range letters {//берём первое значение, которое будем сравнивать
		for d, s := range letters {//берём второе значение, которое будем сравнивать
			if v == s {
				if i != d {//чтобы не сравнивать с самим собой
					fmt.Println("Coincidence " + v + " и " + s)
					notUnique = true
					break
				}
			} else {
				fmt.Println("Not a coincidence " + v + " и " + s)
			}
		}
		if notUnique {
			break
		}
	}

	if notUnique {
		fmt.Println("There are no unique word")
	} else {
		fmt.Println("There are unique word")
	}
}