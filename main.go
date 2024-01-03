package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	inputFileName := "./lang/english.txt"
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer func(inputFile *os.File) {
		_ = inputFile.Close()
	}(inputFile)

	// Filter words by reading line by line from the file and write them to separate files
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		word := scanner.Text()

		// Capitalize the word
		word = strings.ToUpper(word)

		// Skip if the word does not contain alphabetic characters
		if !isValidWord(word) {
			continue
		}

		wordLength := utf8.RuneCountInString(word)
		outputFileName := fmt.Sprintf("%d_letter_words.txt", wordLength)

		// Open the file (or create it if available)
		outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening the file (%s): %v\n", outputFileName, err)
			return
		}

		// Write the word in the file
		_, err = fmt.Fprintf(outputFile, "%s\n", word)
		if err != nil {
			fmt.Printf("Error writing to file (%s): %v\n", outputFileName, err)
			return
		}

		// Close the file now
		_ = outputFile.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading line from file:", err)
		return
	}

	fmt.Println("Process completed.")
}

func isValidWord(word string) bool {
	// Check if there are only letters in the word
	validPattern := regexp.MustCompile(`^[A-ZÇĞİŞÖÜ]+$`)
	return validPattern.MatchString(word)
}
