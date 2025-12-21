package devops

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ExpandVariables expands variables and environment variables in a string.
func ExpandVariables(s string) string {
	for _, envVar := range os.Environ() {
		kv := strings.SplitN(envVar, "=", 2)
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			s = strings.ReplaceAll(s, "${"+k+"}", v)
			s = strings.ReplaceAll(s, "$"+k, v)
		}
	}
	return s
}

// JoinPaths joins a list of paths with the correct OS-specific separator.
func JoinPaths(paths ...string) string {
	return filepath.Join(paths...)
}

// GetExecutablePath returns the path to the current executable.
func GetExecutablePath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("failed to get executable path")
	}
	return filepath.Dir(filename)
}

// ReadFile reads the contents of a file and returns its contents as a string.
func ReadFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// WriteFile writes a string to a file.
func WriteFile(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0o644)
}

// RetryFunc is a function that takes a function and a delay between retries.
type RetryFunc func() error

// Retry executes a function with retries.
func Retry(f RetryFunc, delay float64, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		if err := f(); err == nil {
			return nil
		}
		if i < maxRetries-1 {
			log.Printf("failed, retrying in %fs...\n", delay)
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
	return errors.New("failed after " + fmt.Sprintf("%d retries", maxRetries))
}