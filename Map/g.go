package main

import (
	"fmt"
	"sort"
)

func Exercise4_GroupByGrade(students map[string]int) map[int][]string {

	grouped := make(map[int][]string)

	for key, value := range students {
		grouped[value] = append(grouped[value], key)
	}
	for key := range grouped {
		sort.Strings(grouped[key])
	}
	return grouped
}

// func findMissing(numbers []int, n int) []int {
//     // Hint: Use a map to track which numbers exist
// 	numbers := []int{1, 2, 4, 5, 7}
// 	missing := make(map[int]int)

// }

func main() {
	students := map[string]int{
		"be":  8,
		"d":   2,
		"a":   1,
		"sdl": 5,
	}
	fmt.Println(Exercise4_GroupByGrade(students))
}
