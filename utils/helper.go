package utils

import (
	"fmt"
	"strings"
)

func ParseURL(path string) string {
	fmt.Println("Called ParseURL")
	data := strings.Split(path, "~")
	for i := range data {
		fmt.Println(data[i])
	}
	return path
}
