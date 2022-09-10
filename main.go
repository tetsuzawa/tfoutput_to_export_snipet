package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// IP_X = "xxx.xxx.xxx.xxx"
		rawLine := scanner.Text()
		line := strings.Replace(rawLine, " ", "", -1)
		line = strings.Replace(line, "\"", "", -1)
		line = fmt.Sprintf("export %s", line)
		fmt.Println(line)
	}
}
