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

	for d, v := range letters {
		for i := 1; i < len(letters); i++ {
			if d != i{
				if v == letters[i] {

					fmt.Println("Coincidence " + v + " и " + letters[i])
					notUnique = true
					break

				} else {fmt.Println("Not a coincidence " + v + " и " + letters[i])}//можно ли оставлять так для компактности?
			}
		}
		if notUnique {
			fmt.Println("There are no unique word")
			return
		}
	}
	fmt.Println("There are unique word")
	return

}
