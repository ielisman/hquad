package cli

import (
	"fmt"
	"bufio"
	"os"
)

func PressEnter(msg string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%v", msg)

	scanner.Scan()
}