package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func to_int_array(arr []string) []int {
    var intarr []int
    for _, s := range arr {
        i, _ := strconv.Atoi(s)
        intarr = append(intarr, i)
    }
    return intarr
}

func diff_arr(nums []int) []int {
    var newarr []int
    for i := 1; i < len(nums); i++ {
        newarr = append(newarr, nums[i] - nums[i-1])
    }
    return newarr
}

func all_zeros(nums []int) bool {
    for _, n := range nums {
        if n != 0 {
            return false
        }
    }
    return true
}

func find_next_val(nums []int) int {
    if all_zeros(nums) {
        return 0
    }
    return nums[len(nums)-1] + find_next_val(diff_arr(nums))
}

func find_prev_val(nums []int) int {
    if all_zeros(nums) {
        return 0
    }
    return nums[0] - find_prev_val(diff_arr(nums))
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    next_val := 0
    prev_val := 0
    for scanner.Scan() {
        line := scanner.Text()
        nums := to_int_array(strings.Fields(line))
        next_val += find_next_val(nums)
        prev_val += find_prev_val(nums)
    }
    fmt.Println("Result 1:", next_val)
    fmt.Println("Result 2:", prev_val)
}
