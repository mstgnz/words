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

	classification("./lang/english.txt", "alphabet")

	//checkWords("./lang/english.txt", "./lang/new_english.txt")
	//checkWords("./lang/turkish.txt", "./lang/new_turkish.txt")
}

func classification(fileName, kind string) {

	inputFile, err := os.Open(fileName)
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

		var outputFileName string
		if kind == "alphabet" {
			// Get the first letter of the word
			firstLetter, _ := utf8.DecodeRuneInString(word)
			outputFileName = fmt.Sprintf("lang/%c_letter_words.txt", firstLetter)
		} else {
			wordLength := utf8.RuneCountInString(word)
			outputFileName = fmt.Sprintf("lang/%d_letter_words.txt", wordLength)
		}

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

func checkWords(inputFileName, outputFileName string) {

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer func(inputFile *os.File) {
		_ = inputFile.Close()
	}(inputFile)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		word := scanner.Text()

		// Capitalize the word
		word = strings.ToUpper(word)

		// Skip if the word does not contain alphabetic characters
		if isValidWord(word) {
			// Write the word in the new file
			_, err := fmt.Fprintf(outputFile, "%s\n", word)
			if err != nil {
				fmt.Printf("Error writing to file (%s): %v\n", outputFileName, err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading line from file:", err)
		return
	}
}

func isValidWord(word string) bool {
	// Check if there are only letters in the word
	validPattern := regexp.MustCompile(`^[A-ZÇĞİŞÖÜ]+$`)
	return validPattern.MatchString(word)
}
