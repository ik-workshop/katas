package kata

import (
	"errors"
	"fmt"
)

func TwoSum(numbers []int, target int) ([2]int, error) {
	var result [2]int

	if len(numbers) <= 1 {
		fmt.Println(result)
		return result, errors.New("input slice cannot be empty")
	}

	return result, nil
}
