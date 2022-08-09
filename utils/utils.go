package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadToken() (string, error) {
	data, err := os.ReadFile("./token")

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func GetIntChoice(message string) ([]int, error) {
	var choices []int

	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	choice_string, _ := reader.ReadString('\n')

	choice_string = strings.TrimSpace(choice_string)

	arr := strings.Split(choice_string, " ")

	for _, s := range arr {
		i, err := strconv.Atoi(s)

		if err != nil {
			return choices, err
		}

		choices = append(choices, i)
	}

	return choices, nil
}

func GetString(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	choice_string, _ := reader.ReadString('\n')
	choice_string = strings.TrimSpace(choice_string)

	return choice_string
}

func GetConfirm(message string, positive bool) bool {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	choice_string, _ := reader.ReadString('\n')

	choice_string = strings.TrimSpace(choice_string)

	if strings.ToLower(choice_string) == "y" {
		return positive
	}

	return !positive
}

func PrettyPrintJSON(text string) string {
	var buffer bytes.Buffer

	err := json.Indent(&buffer, []byte(text), "", "\t")
	if err != nil {
		return ""
	}

	return buffer.String()
}

func PrettyPrintJSONBytes(text []byte) string {
	var buffer bytes.Buffer

	err := json.Indent(&buffer, text, "", "\t")
	if err != nil {
		return ""
	}

	return buffer.String()
}
