package util

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// PromptForInt read command line input
func PromptForInt(in io.Reader, out io.Writer, prompt string, defaultValue int) (int, error) {
	PrettyPrintf(out, "=> %s [%d]: ", prompt, defaultValue)
	s := bufio.NewScanner(in)
	// Scan the first token
	s.Scan()
	if s.Err() != nil {
		return defaultValue, fmt.Errorf("error reading number: %v", s.Err())
	}
	ans := s.Text()
	if ans == "" {
		return defaultValue, nil
	}
	// Convert input into integer
	i, err := strconv.Atoi(ans)
	if err != nil {
		return defaultValue, fmt.Errorf("%q is not a number", ans)
	}
	return i, nil
}

// CreateDir check if directory exists and create it
func CreateDir(dir string, perm os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, perm)
		if err != nil {
			return fmt.Errorf("error creating destination dir: %v", err)
		}
	}

	return nil
}

// Base64String read file and return base64 string
func Base64String(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	fileEncoded := base64.StdEncoding.EncodeToString(file)

	return fileEncoded, nil
}
