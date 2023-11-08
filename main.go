package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/manifoldco/promptui"
)

var GESTURES = []string{"Rock", "Paper", "Scissors"}

func main() {
	displayResults := func(playerGesture string, computerGesture string, didWin bool) {
		if didWin {
			fmt.Printf("You won this round!\n You played %q and the Computer played %q\n\n", playerGesture, computerGesture)
		} else {
			fmt.Printf("Computer won this round!\n Computer played %q and you played %q\n\n", computerGesture, playerGesture)
		}
	}

	for {
		var playerWins, computerWins, ties int = 0, 0, 0
		validate := func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		}

		roundsPrompt := promptui.Prompt{
			Label:    "How many rounds do you want to play?",
			Validate: validate,
		}

		input, inputErr := roundsPrompt.Run()
		if inputErr != nil {
			fmt.Printf("Prompt failed %v\n", inputErr)
			break
		}

		rounds, roundsErr := strconv.ParseInt(input, 10, 0)
		if roundsErr != nil {
			fmt.Printf("Parsing failure %v\n", roundsErr)
			break
		}

		for rounds > 0 {
			fmt.Printf("Round %d\n", computerWins+playerWins+ties+1)
			computerGesture := GESTURES[rand.Intn(len(GESTURES)-1)]

			gesturePrompt := promptui.Select{
				Label: "What gesture do you want to play?",
				Items: GESTURES,
			}

			_, result, err := gesturePrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				break
			}

			switch result {
			case "Rock":
				if computerGesture == "Paper" {
					computerWins++
					displayResults(result, computerGesture, false)
				} else if computerGesture == "Scissors" {
					playerWins++
					displayResults(result, computerGesture, true)
				} else {
					fmt.Printf("It's a tie!\n You both played %q\n\n", result)
					ties++
				}
			case "Paper":
				if computerGesture == "Scissors" {
					computerWins++
					displayResults(result, computerGesture, false)
				} else if computerGesture == "Rock" {
					playerWins++
					displayResults(result, computerGesture, true)
				} else {
					fmt.Printf("It's a tie!\n You both played %q\n\n", result)
					ties++
				}
			case "Scissors":
				if computerGesture == "Rock" {
					computerWins++
					displayResults(result, computerGesture, false)
				} else if computerGesture == "Paper" {
					playerWins++
					displayResults(result, computerGesture, true)
				} else {
					fmt.Printf("It's a tie!\n You both played %q\n\n", result)
					ties++
				}
			}
			rounds -= 1
		}

		if playerWins > computerWins {
			fmt.Printf("You won the game! You beaten the computer by %d wins\n", playerWins-computerWins)
		} else if computerWins > playerWins {
			fmt.Printf("You lost the game. The computer beaten you by %d wins\n", computerWins-playerWins)
		} else {
			fmt.Printf("Tie! You and the computer both got %d wins each\n", ties)
		}

		playAgainPrompt := promptui.Select{
			Label: "Would you like to play again?",
			Items: []string{"Yes", "No"},
		}

		_, result, err := playAgainPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			break
		}
		if result == "No" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
