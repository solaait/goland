package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var choice int

	fmt.Println(" ")
	fmt.Println("   " +
		"enter your choice:\n" + "     " + "1:to-----multiplication table\n" + "     " + "2:to-----diamond drawing\n" + "     " + "3:to-----guessing game")

	fmt.Scanln(&choice)

	switch choice {

	case 1:

		multiplication()

	case 2:

		diamond()

	case 3:
		guessing()

	default:

		fmt.Println("you should enter 1,2 or 3")

	}

}

func multiplication() {

	fmt.Print("---multiplication Table------\n\n")

	for i := 1; i <= 12; i++ {

		for j := 1; j <= 12; j++ {

			switch {
			case i*j <= 9:
				fmt.Printf("     %d", i*j)

			case i*j <= 99:
				fmt.Printf("    %d", i*j)

			case i*j >= 100:
				fmt.Printf("   %d", i*j)

			}

		}
		fmt.Println("")

	}
}

///// question number 2

func diamond() {

	diagonal := 5

	index := diagonal - 1

	for i := 1; i <= diagonal; i++ {
		for j := 1; j <= index; j++ {
			fmt.Print(" ")
		}
		index--
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	index = 1
	for i := 1; i <= diagonal-1; i++ {
		for j := 1; j <= index; j++ {
			fmt.Print(" ")
		}
		index++
		for j := 1; j <= 2*(diagonal-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

///////   question number 3

func guessing() {
	secretNumber := secrtNum(1, 20)
	//numGuess := 1
	var guess int

	fmt.Println("      " + "*****************" + "     ")

	fmt.Println("******guess  from 1 to 20 :******")
	fmt.Println("******you have 5 chances:  ******")

	fmt.Println("      " + "*****************" + "     ")

	numGuess := 0
	for numGuess < 5 {

		var arr []int
		fmt.Println("enter your guess :")

		fmt.Scanln(&guess)

		collection := append(arr, guess)

		for index := 0; index < len(collection); index++ {

			if collection[index] == guess {

				//numGuess--

			}

			numGuess++

		}

		if guess < secretNumber {
			fmt.Println("guess is smaller than secret number")

		}

		if guess > secretNumber {
			fmt.Println("guess is greater than secret number")

		}

		if guess == secretNumber {

			fmt.Println("you get the number after ", numGuess, "tries")
			os.Exit(0)
		}

	}

	if guess != secretNumber {

		fmt.Println("you loss!!: the secretNumber was ", secretNumber)

	}

}
func secrtNum(min, max int) int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min

}
