package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	input_names, _ := os.ReadFile("names.txt")
	cleaned_names := extract_names(string(input_names))

	fmt.Println("cleaned_names:", cleaned_names)
	fmt.Println(find_sum_of_letters_in_name("ZZZZ"))
	fmt.Println()

	// name1 := "SUSANNA"
	// name2 := "NATHAN"
	// result := compare_names(name1, name2)
	// fmt.Println(name1, "belongs", result, name2)

	// trying something new

	alphabetical_list := []string{"AAAAAAAA", "ZZZZZZZZ"}

	for _, input_name := range cleaned_names {

		for i := range alphabetical_list {

			result1 := compare_names(alphabetical_list[i], input_name)
			// fmt.Println(existing_name, "belongs", result1, input_name)

			if result1 == "before" {
				result2 := compare_names(alphabetical_list[i+1], input_name)
				// fmt.Println(existing_name, "belongs", result2, input_name)

				if result2 == "after" {
					// fmt.Println(input_name, "should go between", alphabetical_list[i], "and", alphabetical_list[i+1])
					// fmt.Println(alphabetical_list, "i:", i)
					alphabetical_list = append(alphabetical_list, "")
					// fmt.Println(alphabetical_list)
					copy(alphabetical_list[(i+1):], alphabetical_list[i:])
					// fmt.Println(alphabetical_list)
					alphabetical_list[i+1] = input_name
					// 	fmt.Println(alphabetical_list)
					break
				}
			}
		}
	}

	fmt.Println(alphabetical_list)
}

func extract_names(text string) []string { // takes the input names from a .txt file and returns a slice with all of the 'cleaned' names
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

func compare_names(existing_name string, input_name string) string {

	if existing_name == input_name {
		fmt.Println("error: name1 and name2 are the same")
	}

	var determination string = "undecided"
	shorter_name := input_name // set the input_name as shorter, unless proven otherwise
	letter_to_number := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7,
		"H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17,
		"R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}

	if len(existing_name) < len(input_name) {
		shorter_name = existing_name
	}

	for i := range shorter_name {
		existing_name_letter := letter_to_number[string(existing_name[i])]
		input_name_letter := letter_to_number[string(input_name[i])]
		if existing_name_letter > input_name_letter {
			determination = "after"
			break
		} else if existing_name_letter < input_name_letter {
			determination = "before"
			break
		} else {
			// continue to the next letter to check
		}
	}

	if determination == "undecided" {
		if existing_name == shorter_name {
			determination = "before"
		} else if input_name == shorter_name {
			determination = "after"
		}
	}
	return determination
}
