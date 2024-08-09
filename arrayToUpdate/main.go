package main

import (
	"os"
	"strings"
	// "fmt"
)


func main() {

	data, err := os.ReadFile("data.txt")
	if err != nil {
		os.Exit(1)
	}

	data2, err2 := os.ReadFile("data2.txt")
	if err2 != nil {
		os.Exit(1)
	}

	dataArray := strings.Split(string(data), "\n")
	data2Array := strings.Split(string(data2), "\n")

	var res string
	for i := range dataArray {


		res += "UPDATE item SET mdm_nr = '" + strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(data2Array[i]), "\r\n", "\n"), "\n", "")) + "' WHERE name = '" + strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(dataArray[i]), "\r\n", "\n"), "\n", "")) + "'; \n"
	}
	// res = res[0:len(res)-2]

	os.WriteFile("./output.txt", []byte(res), 0644)
}
