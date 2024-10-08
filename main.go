package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int, charset string) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func GeneratePassword(length int, includeUppercase, includeNumbers, includeSymbols bool) string {
	charset := "abcdefghijklmnopqrstuvwxyz"

	if includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSymbols {
		charset += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}

	return GenerateRandomString(length, charset)
}

func main() {
	length := flag.Int("length", 12, "Set the length of the password")
	includeUppercase := flag.Bool("uppercase", true, "Include uppercase letters in the password")
	includeNumbers := flag.Bool("numbers", true, "Include numbers in the password")
	includeSymbols := flag.Bool("symbols", true, "Include symbols in the password")
	copyToClipboard := flag.Bool("copy", false, "Copy the generated password to the clipboard")

	flag.Parse()

	password := GeneratePassword(*length, *includeUppercase, *includeNumbers, *includeSymbols)

	fmt.Printf("Generated Password: %s\n", password)

	if *copyToClipboard {
		if err := CopyToClipboard(password); err != nil {
			fmt.Println("Error copying to clipboard:", err)
		} else {
			fmt.Println("Password copied to clipboard!")
		}
	}
}

// todo implement clipboard copy
func CopyToClipboard(text string) error {
	return nil
}
