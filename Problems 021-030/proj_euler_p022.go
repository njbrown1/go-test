package main

import "fmt"
import "regexp"
import "os"

func main() {
	input_names, _ := os.ReadFile("name_testing.txt")
	cleaned_names := extract_names(string(input_names))
	fmt.Println(cleaned_names)
	fmt.Println(find_sum_of_letters_in_name("COLIN"))
}

func extract_names(text string) []string { // takes the input names and returns a slice with all of the 'cleaned' names
	re := regexp.MustCompile(`\s*"([^"]+)"\s*`)
	matches := re.FindAllStringSubmatch(text, -1) // -1 finds all matches, apparently

	var names []string
	for _, match := range matches {
		name := match[1] // extract the match, apparently?
		names = append(names, name)
	}
	return names
}

func find_sum_of_letters_in_name(name string) int {
	var alphabet string = ":ABCDEFGHIJKLMNOPQRSTUVWXYZ" // ':' is just an arbitrary start character
	sum_of_letters := 0

	for i := range name {
		nth_letter := name[i]
		for letter := 1; letter <= 26; letter++ {
			if alphabet[letter] == nth_letter {
				sum_of_letters += letter
				break
			}
		}
	}
	return sum_of_letters
}
