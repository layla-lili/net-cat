package colortest

import "testing"

// Define ANSI escape codes for colors
const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
)

// Define custom logging functions with color prefixes
func LogError(t *testing.T, msg string) {
    t.Errorf(Red + "Error: " + msg + Reset)
}

func LogSuccess(t *testing.T, msg string) {
    t.Logf(Green + "Success: " + msg + Reset)
}

func LogInfo(t *testing.T, msg string) {
    t.Logf(Yellow + "Info: " + msg + Reset)
}