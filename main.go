// Coldtrail - a number-guessing duel in Go.
//
// A number between 1 and 100 is hidden. You have 7 guesses per round.
// After each guess you are told how close you are (the "temperature")
// and whether to go higher or lower. Three rounds; your score is the
// total of the guesses you had left, converted to points.
// Usage: coldtrail [--seed N]
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func temperature(guess, target int) string {
	d := guess - target
	if d < 0 {
		d = -d
	}
	switch {
	case d == 0:
		return "ON THE TRAIL"
	case d <= 2:
		return "burning hot"
	case d <= 5:
		return "hot"
	case d <= 10:
		return "warm"
	case d <= 20:
		return "cool"
	case d <= 35:
		return "cold"
	default:
		return "freezing"
	}
}

func main() {
	seed := int64(20261125)
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "--seed" && i+1 < len(args) {
			if v, err := strconv.ParseInt(args[i+1], 10, 64); err == nil {
				seed = v
			}
			i++
		}
	}
	rng := rand.New(rand.NewSource(seed))
	in := bufio.NewReader(os.Stdin)

	const rounds = 3
	const maxGuesses = 7
	total := 0

	fmt.Println("=== COLDTRAIL ===")
	fmt.Println("A number from 1 to 100 is hidden on the trail.")
	fmt.Printf("%d rounds, %d guesses each. Follow the temperature.\n\n", rounds, maxGuesses)

	for r := 1; r <= rounds; r++ {
		target := rng.Intn(100) + 1
		fmt.Printf("--- Round %d/%d ---\n", r, rounds)
		found := false
		for g := 1; g <= maxGuesses; g++ {
			fmt.Printf("Guess %d/%d: ", g, maxGuesses)
			line, err := in.ReadString('\n')
			if err != nil && strings.TrimSpace(line) == "" {
				fmt.Println()
				goto done
			}
			guess, cerr := strconv.Atoi(strings.TrimSpace(line))
			if cerr != nil || guess < 1 || guess > 100 {
				fmt.Println("Enter a whole number from 1 to 100.")
				g--
				continue
			}
			if guess == target {
				pts := (maxGuesses - g + 1) * 20
				total += pts
				fmt.Printf("Got it in %d! +%d pts (total %d)\n\n", g, pts, total)
				found = true
				break
			}
			dir := "higher"
			if guess > target {
				dir = "lower"
			}
			fmt.Printf("%s - go %s.\n", temperature(guess, target), dir)
		}
		if !found {
			fmt.Printf("Out of guesses. The number was %d.\n\n", target)
		}
	}

done:
	fmt.Println("=== GAME OVER ===")
	fmt.Printf("Final score: %d\n", total)
	switch {
	case total >= 300:
		fmt.Println("Rating: Trailblazer")
	case total >= 180:
		fmt.Println("Rating: Tracker")
	case total >= 60:
		fmt.Println("Rating: Wanderer")
	default:
		fmt.Println("Rating: Lost")
	}
}
