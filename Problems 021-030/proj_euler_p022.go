package main

import (
	"fmt"
	"os"
	"regexp"
)

// A quick high-level overview of the code:
// 1. the extract_names function takes all of the names from names.txt and puts them into the cleaned_names slice (to get
// rid of the quotation marks, commas, and whitespace). I had to learn Regex to solve that problem effectively!
// 2. the compare_names function just determines the alphabeticized order of existing_name and input_name.
// 3. the find_sum_of_letters_in_name function: if the letters in the name COLIN were turned into numbers (ex. C = 3, O = 15, L = 12),
// the function would find the sum of those numbers. so find_sum_of_letters_in_name("COLIN") = 53.
// 4. here's what main() does: for EVERY name in the cleaned_names slice, it tries to find the two names that already EXIST in
// the alphabetical_list slice that correctly 'sandwich' the input_name (ie. the input_name fits between those two names). once it's found
// those two names, it INSERTS the input_name at the correct position in the alphabetical_list slice (using copying techniques). note: I added
// "AAAAAAAA" and "ZZZZZZZZ" to begin the alphabetical list, so I didn't have to modify the 'for' loop for the first two names (because
// those first two names would not have bread to sandwich themselves between if I didn't add "AAAAAAAA" and "ZZZZZZZZ").
// 5. after the names are all alphabeticized correctly in the alphabetical_list slice, each name's "name score" is determined by MULTIPLYING
// the output from the find_sum_of_letters_in_name function and the numbered position of the name within the slice. ex. COLIN would obtain a
// score of 938 x 53 = 49714. and all of the name scores are added together – and I was careful to ensure that the name scores for "AAAAAAAA"
// and "ZZZZZZZZ" were NOT added to the final sum_of_name_scores.

func main() {

	input_names, _ := os.ReadFile("names.txt")
	cleaned_names := extract_names(string(input_names))
	fmt.Println("cleaned_names:", cleaned_names)
	alphabetical_list := []string{"AAAAAAAA", "ZZZZZZZZ"}

	for _, input_name := range cleaned_names {
		for i := range alphabetical_list {
			result1 := compare_names(alphabetical_list[i], input_name)
			if result1 == "before" {
				result2 := compare_names(alphabetical_list[i+1], input_name)
				if result2 == "after" {
					alphabetical_list = append(alphabetical_list, "")
					copy(alphabetical_list[(i+1):], alphabetical_list[i:])
					alphabetical_list[i+1] = input_name
					break
				}
			}
		}
	}
	sum_of_name_scores := 0
	for i := 1; i < (len(alphabetical_list) - 1); i++ {
		sum_of_name_scores += (find_sum_of_letters_in_name(alphabetical_list[i]) * i)
	}
	fmt.Println(sum_of_name_scores) // solved! answer: 871198282
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

	if existing_name == input_name { // all of the names in names.txt are distinct, thankfully.
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
			// continue to the next letter
		}
	}

	if determination == "undecided" { // because the English alphabetical conventions say that ANNA would fall before ANNABELLA.
		if existing_name == shorter_name {
			determination = "before"
		} else if input_name == shorter_name {
			determination = "after"
		}
	}
	return determination
}
