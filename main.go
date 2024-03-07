package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func magicEightBall() {
	messages := []string{"It is certain", "Very doubtful", "Ask again later",
		"Without a doubt", "As I see it, yes", "Don't count on it", "Cannot predict now",
		"Outlook good", "Most likely", "Yes", "My source says no", "Reply hazy, try again",
		"It is decidedly so", "Better not tell you now", "Outlook not so good", "My reply is no",
		"You may rely on it", "Definitely yes", "Signs point to yes"}

	for {
		randNum := randomInt(0, len(messages))

		fmt.Println("* * * Magic Eight Ball of Mystery * * *")
		fmt.Printf("Please ask your question (enter 'q' to quit): ")

		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Remove newline character from the input
		answer = strings.TrimSpace(answer)

		if answer == "q" {
			fmt.Println("Thank you for finding your fortune!")
			return
		}

		fmt.Print("... ")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("... ")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("...\n")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(messages[randNum], "\n")
	}
}

func main() {
	magicEightBall()
}
