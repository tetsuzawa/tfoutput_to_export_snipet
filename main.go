package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	builder := &strings.Builder{}
	for scanner.Scan() {
		// IP_X = "xxx.xxx.xxx.xxx"
		rawLine := scanner.Text()
		line := strings.Replace(rawLine, " ", "", -1)
		line = strings.Replace(line, "\"", "", -1)
		builder.WriteString(fmt.Sprintf("export %s", line))
		builder.WriteString(" && \\\n")
	}
	res := strings.TrimSuffix(builder.String(), " && \\\n")
	fmt.Println(res)
}
