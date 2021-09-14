package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//Add a second dice
//Limit maximum number of tries
func main() {
	fmt.Println("Roll the dice and try to guess the result! Or press 'q' to quit!")
	var input string
	result := rand.Intn(6) + 1

	for {
		fmt.Print("Your guess: ")
		fmt.Scan(&input)

		if input == "q" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid number please!")
			continue
		}

		if num == result {
			fmt.Println("Congrats! You guessed it right! The result is: ", result)
		}

	}

}
