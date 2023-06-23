package main

import "fmt"
import "regexp"
import "os"

func main() {
	fmt.Println("hello world")

	input_names, _ := os.ReadFile("name_testing.txt")
	cleaned_names := extract_names(string(input_names))
	fmt.Println(cleaned_names)

}

func extract_names(text string) []string {
	re := regexp.MustCompile(`\s*"([^"]+)"\s*`)
	matches := re.FindAllStringSubmatch(text, -1) // -1 finds all matches, apparently

	var names []string
	for _, match := range matches {
		name := match[1] // extract the match, apparently?
		names = append(names, name)
	}
	return names
}
