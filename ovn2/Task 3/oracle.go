// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

func answerQuestion(q string, answer chan string) {
	time.Sleep(3 * time.Second)
	replies := []string{"That's not accurate",
		"Please, don't ask that!",
		"Forget it - I won't tell!",
		"Maybe...",
		"That's reverse psychology. I know what you're up to!",
		"Stop asking so many questions, will ya?",
		"Forget about it.",
		"Where's Joe?"}
	answer <- replies[rand.Intn(len(replies))]
}

func receiveQuestion(questions <-chan string, answer chan string) {
	for q := range questions {
		go answerQuestion(q, answer)
	}
}

func generateProphecies(answer chan string) {
	for {
		time.Sleep(10 * time.Second)
		prophecy("", answer)
	}
}

func receivePropheciesAndAnswers(answer chan string) {
	for a := range answer {
		for _, char := range strings.Split(a, "") {
			fmt.Print(char)
			time.Sleep(time.Duration(35+rand.Intn(265)) * time.Millisecond)
		}
		fmt.Println()
		fmt.Print("> ")
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answer := make(chan string)
	// TODO: Answer questions.
	go receiveQuestion(questions, answer)
	// TODO: Make prophecies.
	go generateProphecies(answer)
	// TODO: Print answers.
	go receivePropheciesAndAnswers(answer)
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"You will die some day.",
		"Judgement Day is close.",
		"You will be involved in an accident in the next 3 decades.",
		"Some day, you will have children; one boy and one daughter.",
		"If you end up not having a daughter, you will end up having two boys.",
		"I know, therefore I tell. And so will you, once you know.",
	}
	answer <- "\n" + prompt + longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
