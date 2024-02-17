package main

import "fmt"

const (
    RESET = "\x1b[0m"
    BOLD = "\x1b[1m"
    UNDERLINE = "\x1b[4m"
    BLACK = "\x1b[30m"
    RED = "\x1b[31m"
    GREEN = "\x1b[32m"
    YELLOW = "\x1b[33m"
    BLUE = "\x1b[34m"
    MAGENTA = "\x1b[35m"
    CYAN = "\x1b[36m"
    WHITE = "\x1b[37m"
    CLEAR = "\x1b[2J\x1b[H"  // clear screen and move cursor to 0,0
)

func Black(s string) string {
    return fmt.Sprintf("%s%s%s", "\x1b[30m", s, RESET)
}

func Red(s string) string {
    return fmt.Sprintf("%s%s%s", RED, s, RESET)
}

func Green(s string) string {
    return fmt.Sprintf("%s%s%s", GREEN, s, RESET)
}

func Yellow(s string) string {
    return fmt.Sprintf("%s%s%s", YELLOW, s, RESET)
}

func Blue(s string) string {
    return fmt.Sprintf("%s%s%s", BLUE, s, RESET)
}

func Magenta(s string) string {
    return fmt.Sprintf("%s%s%s", MAGENTA, s, RESET)
}

func Cyan(s string) string {
    return fmt.Sprintf("%s%s%s", CYAN, s, RESET)
}

func White(s string) string {
    return fmt.Sprintf("%s%s%s", WHITE, s, RESET)
}

func Bold(s string) string {
    return fmt.Sprintf("%s%s%s", BOLD, s, RESET)
}

func Underline(s string) string {
    return fmt.Sprintf("%s%s%s", UNDERLINE, s, RESET)
}
