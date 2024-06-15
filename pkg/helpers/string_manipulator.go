package helpers

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/mail"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// validation email format
// example here: https://go.dev/play/p/j4B4v01Qolw
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

// Lower case a string, ex; HELLO WORLD => hello world
func ToLower(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

// Capitalize each word in a string, ex; hello world => Hello World
func ToTitle(input string) string {
	var (
		output []rune
		isWord = true
	)

	for _, val := range input {
		if isWord && unicode.IsLetter(val) { //check if character is a letter convert the first character to upper case
			output = append(output, unicode.ToUpper(val))
			isWord = false
		} else if !unicode.IsLetter(val) {
			isWord = true
			output = append(output, val)
		} else {
			output = append(output, val)
		}
	}

	return strings.TrimSpace(string(output))
}

// Sentence case in a string, ex; hello world => Hello world
func ToSentenceCase(input string) string {
	input = ToLower(input)
	if len(input) <= 0 {
		return ""
	}

	temp := strings.Split(input, " ")
	temp[0] = ToTitle(temp[0])
	return strings.Join(temp, " ")
}

// Upper case a string, ex; hello world => HELLO WORLD
func ToUpper(input string) string {
	return strings.TrimSpace(strings.ToUpper(input))
}

// GenerateRandomNumber generates a random number of the specified length (length int) and returns it as a string (string).
func GenerateRandomNumber(length int) string {
	max := int(math.Pow10(length))
	num := rand.Intn(max)
	return fmt.Sprintf("%0*d", length, num)
}
