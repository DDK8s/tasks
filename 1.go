package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var notUnique bool
	newLetters := []string{""}
	fmt.Printf("Enter a word: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	word := sc.Text()

	letters := strings.Split(word, "")
	fmt.Println(letters)

	newLetters = append(newLetters, letters...)


	for i, v := range letters {
		for d, s := range letters {

			if d != 0 && s == letters[d-1] {
				if i != d {

					if v == s {
						fmt.Println("Coincidence " + v + " и " + s)
						notUnique = true
						break
					} else {
						fmt.Println("Not a coincidence " + v + " и " + s)
						if d != 0 && s == letters[d+1] {
							break
						}
					}

				}
			}
		}
		if notUnique {
			fmt.Println("There are no unique word")
			return
		}else {
			fmt.Println("There are unique word")
			return
		}
	}
}
