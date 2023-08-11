package src

import (
	"os"
	"strings"
)

func CommaSeparated(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}

	dataArray := strings.Split(strings.TrimSpace(strings.ReplaceAll(string(data), "\r\n", "\n")), "\n")

	var res string
	for _, datarow := range dataArray {
		res += datarow + ", "
	}
	res = res[0:len(res)-2]

	os.WriteFile("./output.txt", []byte(res), 0644)
}