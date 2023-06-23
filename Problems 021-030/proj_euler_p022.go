package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	input_names, _ := os.ReadFile("name_testing.txt")
	cleaned_names := extract_names(string(input_names))
	fmt.Println(cleaned_names)
	fmt.Println(find_sum_of_letters_in_name("ZZZZ"))
	fmt.Println(compare_names("ANDREW", "ANDREA"))

	// alphabetical_list_of_names := []string{"AAAAAAAA", "ZZZZZZZZ"}

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
		nth_letter := name[i] //
		for letter := 1; letter <= 26; letter++ {
			if alphabet[letter] == nth_letter {
				sum_of_letters += letter
				break
			}
		}
	}
	return sum_of_letters
}

func compare_names(pre_existing_name string, name_to_be_added string) bool {

	var determination bool // determination is 0 if

	letter_to_number := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8,
		"I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19,
		"T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}

	for letter := range name_to_be_added {
		pre_existing_name_LETTER := letter_to_number[string(pre_existing_name[letter])]
		name_to_be_added_LETTER := letter_to_number[string(name_to_be_added[letter])]
		fmt.Println(pre_existing_name_LETTER, name_to_be_added_LETTER)

		if pre_existing_name_LETTER > name_to_be_added_LETTER {
			fmt.Println(name_to_be_added, "belongs before", pre_existing_name)
			determination = true
			break
		} else if pre_existing_name_LETTER < name_to_be_added_LETTER {
			fmt.Println(name_to_be_added, "belongs after", pre_existing_name)
			determination = false
			break
		} else if pre_existing_name_LETTER == name_to_be_added_LETTER {
			fmt.Println("need to continue")
		}
	}
	return determination
}
