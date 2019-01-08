word score
----------------------
package main
 
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)
 
func jumbled(word string) string {
	var jumble string
	for len(word) > 0 {
		i := rand.Intn(len(word))
		jumble += string(word[i])
		word = word[:i] + word[i+1:]
	}
	return jumble
}
 
func init() {
	rand.Seed(time.Now().UnixNano())
}
 
func main() {
	Words := []string{"bash", "bath", "Beach", "black", "blame", "blanket", "blaring", "Branch", "brick", "Brim", "Champion", "chance", "Chapter", "Charge", "cheese", "Chew", "child", "children", "chili", "Chimpanzee", "China", "chore", "claim", "clean", "clock", "Coach", "cradle", "crate", "crisp", "crop", "drama", "Drape", "drink", "drive", "drug", "drum", "Earth", "flake", "flame", "flamingo", "float", "Fourth", "Freeze", "French", "friend", "frog", "Frost", "glaze", "gleam", "Grade", "grain", "grand", "grin", "grip", "Growth", "grunt", "gush", "Health", "inch", "mask", "month", "plank", "plant", "Punch", "rich", "sand", "scan", "score", "Scrapbook", "shadow", "shake", "Shampoo", "Sheep", "shelf", "shell", "shine", "Shore", "Short", "should", "shower", "skate", "Slay", "Slime", "slush", "Smooth", "smuggle", "snake", "snap", "snare", "sneeze", "snow", "snug", "space", "Spark", "special", "spill", "sport", "stale", "stall", "stampede", "starch", "steal", "strange", "stuck", "swell", "swift", "swim", "swing", "Switch", "than", "Thank", "then", "there", "Thick", "Thief", "think", "Throw", "Thumb", "Thursday", "tooth", "trash"}
	fmt.Printf(`
		Welcome to Word Jumble!
		Unscramble the letters and find the word.
		`)
 
	points := 0
	input := bufio.NewScanner(os.Stdin)
 
	for len(Words) > 0 {
		randomNum := rand.Intn(len(Words))
		randomWord := Words[randomNum]
		jumbledWord := jumbled(randomWord)
		fmt.Println("\nThe jumbled word is: ", jumbledWord)
		fmt.Println("Your guess is:", string(randomWord[0]), "....")
		var guess string
		if input.Scan() {
			guess = input.Text()
		}
		if guess == randomWord {
			fmt.Println("Thats Correct. Excellent.!!")
			points += 100
			fmt.Println("Your score:", points)
		} else {
			fmt.Println("Thats not correct")
			fmt.Println("Correct word is:", randomWord)
		}
		if len(Words) > 1 {
			firstWords := Words[:randomNum]
			lastWords := Words[randomNum+1:] //Remove this word from the Word list
			Words = append(firstWords, lastWords...)
		} else {
			break
		}
		fmt.Println("Do you want to play again? (y/n)")
		if input.Scan() {
			play := input.Text()
			if play != "y" {
				break
			}
		}
 
	}
	fmt.Println("\nThanks for playing.")
	fmt.Println("\nYour SCORE:", points)
	fmt.Println("Congratulations..!!!..See you again.")
}
