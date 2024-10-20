package main

import "fmt"

func main() {
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println(formatIP(ip))

	var numbers []int
	var err error
	numbers, err = listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Четные числа: ", numbers)
	}

	str := "aababc"
	result := countCharacters(str)
	for char, count := range result {
		fmt.Printf("Символ '%c' встречается %d раз\n", char, count)
	}

}

func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, fmt.Errorf("левая граница больше правой")
	}
	var evens []int
	for i := start; i < end+1; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

func countCharacters(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	return charCount
}
