package main

// @author abbybrown
// @date 03/06/2024
// @file main.go
// 		This file lets the user get their fortune by taking in the user's input and returning their
//		fortune from a list of randomly generated responses.

// make sure to put import statements in ()
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
*
@param min : int : the minimum index to select from
@param max : int : the maximum index to select from
@return random number between min and max (but can only take one argument, so it's kind of weird
*/
func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

/*
*
Lets the user ask a question as many times as they want and will give a random answer back
*/
func magicEightBall() {
	// create all the messages possible
	messages := []string{"It is certain", "Very doubtful", "Ask again later",
		"Without a doubt", "As I see it, yes", "Don't count on it", "Cannot predict now",
		"Outlook good", "Most likely", "Yes", "My source says no", "Reply hazy, try again",
		"It is decidedly so", "Better not tell you now", "Outlook not so good", "My reply is no",
		"You may rely on it", "Definitely yes", "Signs point to yes"}

	// apparently while loops don't exist, so do a for loop that checks for 'q'
	for {
		randNum := randomInt(0, len(messages))

		// print the menu
		fmt.Println("* * * Magic Eight Ball of Mystery * * *")
		fmt.Printf("Please ask your question (enter 'q' to quit): ")

		// see if there is an error and return that error, then quit
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// split apart the user's string answer
		answer = strings.TrimSpace(answer)

		// if the user entered 'q,' then quit
		if answer == "q" {
			fmt.Println("Thank you for finding your fortune!")
			return
		}

		// do a fun time.Sleep thing to build suspense
		fmt.Print("... ")
		time.Sleep(600 * time.Millisecond)
		fmt.Print("... ")
		time.Sleep(600 * time.Millisecond)
		fmt.Print("...\n")
		time.Sleep(200 * time.Millisecond)
		fmt.Println(messages[randNum], "\n") // I want that new line for style purposes
	}
}

func main() {
	// call our function!
	magicEightBall()
}
