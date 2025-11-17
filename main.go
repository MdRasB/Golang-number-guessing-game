package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")

	var t int
	var totalPoint int
	fmt.Print("How many match you want to play: ")
	fmt.Scan(&t)

	for ;t>0;t-- {

		n, m := 1, 100
		num:= rand.IntN(m - n + 1 ) + n

		gameMassage()
		yourPoint := yourDificulty(num)

		totalPoint += yourPoint
	}

	fmt.Printf("Your final point is : %v", totalPoint);
}

func gameMassage(){
	fmt.Println("\nI am thinking of a number between 1 and 100.")
	fmt.Println("You have to guess the currect number.");

	fmt.Println("Please select the dificulty level:")
	fmt.Println("1.Easy (10 changes)\n2.Medium (5 changes)\n3.Hard (3 changes)")

}

func yourDificulty(num int) int {
	var choice int
	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&choice)
	var result bool
	var point int

	switch (choice){
		case 1:{
			fmt.Println("\nGreat! You have selected the Easy dificulty level.")
			fmt.Println("You have 10 guesses to get the correct number.")
			fmt.Println("Let's start the game!")
			fmt.Println()
			result, point = maingame(num, 10)
			break
		}
		case 2:{
			fmt.Println("\nGreat! You have selected the Medium dificulty level.")
			fmt.Println("You have 5 guesses to get the correct number.")
			fmt.Println("Let's start the game!")
			fmt.Println()
			result, point = maingame(num, 5)
			break
		}
		case 3:{
			fmt.Println("\nGreat! You have selected the Hard dificulty level.")
			fmt.Println("You have 3 guesses to get the correct number.")
			fmt.Println("Let's start the game!")
			fmt.Println()
			result, point = maingame(num, 3)
			break
		}
	}

	if result {
		fmt.Println("\nYou have earned points.")
		fmt.Printf("Your point is %v\n\n", point)
	}else {
		fmt.Println("\nYou Lost...")
		fmt.Printf("The correct number is : %v\n", num);
		fmt.Println("You haven't earn any points.")
		fmt.Println()
	}
	return point
}

func maingame(num, changes int) (bool, int){
	youWin := false
	var guess int
	remainingChanges := 0

	for i:=1;i<=changes;i++{
		fmt.Print("\nEnter your guess: ")
		fmt.Scan(&guess)

		if guess == num {
			fmt.Printf("Congratulations! You guessed the correct number in %v attemps\n", i)
			youWin = true;
			remainingChanges = changes - i +1
			break;
		} else if guess > num {
			fmt.Printf("Incorrect! The number is less than %v\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is greater than %v\n", guess)
		}
	}
	return youWin, remainingChanges
}