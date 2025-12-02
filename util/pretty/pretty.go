package pretty

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Print[T any](a T) {
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}

func Grid[T any](grid [][]T) {
	b, _ := json.Marshal(grid)
	str := strings.ReplaceAll(string(b), "],", "]\n")
	str = str[1:]
	str = str[:len(str)-1]
	fmt.Println(str)

}
