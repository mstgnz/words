/*
   This is a multiline comment.
   It can span multiple lines.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Giriş dosyasını aç
	inputFileName := "turkish.txt"
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Giriş dosyasını açarken hata oluştu:", err)
		return
	}
	defer inputFile.Close()

	// Dosyadan satır okuyarak kelimeleri filtrele ve ayrı dosyalara yaz
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		word := scanner.Text()

		// Kelimeyi büyük harfe çevir
		word = strings.ToUpper(word)

		// Eğer kelime alfabetik karakterler içermiyorsa geç
		if !isValidWord(word) {
			continue
		}

		wordLength := len(word)
		outputFileName := fmt.Sprintf("%d_letter_words.txt", wordLength)

		// Dosyayı aç (veya varsa oluştur)
		outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Dosyayı açarken hata oluştu (%s): %v\n", outputFileName, err)
			return
		}

		// Kelimeyi dosyaya yaz
		_, err = fmt.Fprintf(outputFile, "%s\n", word)
		if err != nil {
			fmt.Printf("Dosyaya yazarken hata oluştu (%s): %v\n", outputFileName, err)
			return
		}

		// Dosyayı hemen kapat
		outputFile.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Dosyadan satır okurken hata oluştu:", err)
		return
	}

	fmt.Println("İşlem tamamlandı.")
}

func isValidWord(word string) bool {
	// Kelime içinde sadece harf içerip içermediğini kontrol et
	validPattern := regexp.MustCompile(`^[A-Z]+$`)
	return validPattern.MatchString(word)
}

