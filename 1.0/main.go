package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Дан список чисел. Определите, сколько в этом списке элементов, которые больше двух
своих соседей и выведите количество таких элементов.
*/

func neighbors(arr []int) {
	count := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	//text, _ := reader.ReadString('\n')
	//text = strings.TrimSpace(text)
	//n, _ := strconv.Atoi(text)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	slice := strings.Split(text, " ")

	var a = make([]int, len(slice))
	for i := 0; i < len(a); i++ {
		a[i], _ = strconv.Atoi(slice[i])
	}

	neighbors(a)
}
