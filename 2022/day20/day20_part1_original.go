// https://adventofcode.com/2022/day/20
// https://github.com/Favo02/advent-of-code

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Number struct {
	value int
	index int
}

func main() {
	numbers := parseInput()
	nums := orderNumbers(numbers, 1)
	n1000, n2000, n3000 := get10002000300(nums)
	fmt.Println(nums[n1000].value + nums[n2000].value + nums[n3000].value)
}

const DECRYPTION_KEY int = 811589153

func parseInput() (nums []Number) {
	scanner := bufio.NewScanner(os.Stdin)
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		nums = append(nums, Number{n, i})
		i++
	}
	return nums
}

func orderNumbers(nums []Number, times int) []Number {

	for i := 0; i < times; i++ {
		indexToMove := 0
		for nMoved := 0; nMoved < len(nums); nMoved++ {
			for i := 0; i < len(nums); i++ {
				if nums[i].index == indexToMove {
					nums = moveNumber(nums, i, nums[i].value)
					indexToMove++
					break
				}
			}
		}
		fmt.Println("ordered", i+1, "times:")
		fmt.Println(printNums(nums))
	}

	return nums
}

func moveNumber(nums []Number, index, value int) []Number {
	if value == 0 {
		return nums
	} else if value > 0 {
		return doForwardSwaps(nums, index, value%len(nums))
	} else {
		return doBackwardSwaps(nums, index, value%len(nums))
	}
}

func doForwardSwaps(nums []Number, index, times int) []Number {
	swap := reflect.Swapper(nums)
	for i := 0; i < times; i++ {
		swap((index+i)%len(nums), (index+i+1)%len(nums))
	}
	return nums
}

func doBackwardSwaps(nums []Number, index, times int) []Number {
	swap := reflect.Swapper(nums)
	for i := 0; i < -times; i++ {
		swap(((index-i)%len(nums)+len(nums))%len(nums), ((index-i-1)%len(nums)+len(nums))%len(nums))
	}
	return nums
}

func printNums(nums []Number) (res string) {
	for _, n := range nums {
		res = fmt.Sprint(res, n.value, " ")
	}
	return res
}

func get10002000300(nums []Number) (int, int, int) {
	var index int
	for i, n := range nums {
		if n.value == 0 {
			index = i
		}
	}
	return (index + 1000) % len(nums), (index + 2000) % len(nums), (index + 3000) % len(nums)
}
