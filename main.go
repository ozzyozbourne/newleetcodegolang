package main

import (
	p "fmt"
	"sort"
)

func main() {
	printName("ozzy", 5)
	printLinearlyFrom1ToN(10)
	printLinearlyFromNTo1(10)
	p.Printf("%d\n", summation(4))
	p.Printf("%d\n", factorialB(4))
	p.Printf("%d\n", factorialTailOp(4))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	reverseRecursively(arr, 0)
	p.Printf("%v\n", arr)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverseRecursively(arr, 0)
	p.Printf("%v\n", arr)

	p.Printf("%v\n", isPalindrome([]byte("osaid"), 0))
	p.Printf("%v\n", isPalindrome([]byte("OohoO"), 0))
	p.Printf("%d\n", fib(3))
	p.Printf("%d\n", fib(4))
	p.Printf("%d\n", fib(5))

	isSubsequence("b", "abc")
	a := [][]int{}
	p.Printf("%v\n", a == nil)
	fourSum([]int{2, 2, 2, 2}, 8)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	set := make(map[[3]int]struct{})
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				set[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	for triplet := range set {
		res = append(res, []int{triplet[0], triplet[1], triplet[2]})
	}
	return res

}

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func printName(name string, num int) {
	if num == 0 {
		return
	}
	p.Printf("%s\n", name)
	printName(name, num-1)
}

func printLinearlyFrom1ToN(n uint) {
	if n == 0 {
		return
	}
	printLinearlyFrom1ToN(n - 1)
	p.Printf("%d\n", n)
}

func printLinearlyFromNTo1(n uint) {
	if n == 0 {
		return
	}
	p.Printf("%d\n", n)
	printLinearlyFromNTo1(n - 1)
}

func summation(n uint) uint {
	return summationOfNTo1(n, 0)
}

func summationOfNTo1(n uint, sum uint) uint {
	if n == 0 {
		return sum
	}
	return summationOfNTo1(n-1, sum+n)
}

func factorialTailOp(n uint) uint {
	return factorialA(n, 1)
}

func factorialA(n uint, sum uint) uint {
	if n == 1 {
		return sum
	}
	return factorialA(n-1, sum*n)
}

func factorialB(n uint) uint {
	if n == 1 {
		return 1
	}
	return n * factorialB(n-1)
}

func reverseRecursively(arr []int, i int) {
	if i >= len(arr)-1-i {
		return
	}
	temp := arr[i]
	arr[i] = arr[len(arr)-1-i]
	arr[len(arr)-1-i] = temp
	reverseRecursively(arr, i+1)
}

func isPalindrome(v []byte, i int) bool {
	if i == len(v) {
		return true
	}
	if v[i] != v[len(v)-1-i] {
		return false
	}
	return isPalindrome(v, i+1)
}

func isSubsequence(s string, t string) bool {
	l, c, d := 0, len(s), len(t)
	for i := 0; l < c && i < d; i++ {
		if s[l] == t[i] {
			l++
		}
	}
	return l == c
}

func threeSumOptimal(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if sum < 0 {
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	var res []int
	for i, v := range nums {
		if _, ok := dict[target-v]; ok {
			res = append(res, i, dict[target-v])
			break
		} else {
			dict[v] = i
		}
	}
	return res
}

func twoSumtwo(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] < target {
			l += 1
		} else if nums[l]+nums[r] > target {
			r -= 1
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	return kSum(4, 0, int64(target), nums, [][]int{}, []int{})
}

func kSum(k int, start int, target int64, nums []int, res [][]int, tuple []int) [][]int {
	if k > 2 {
		for i := start; i < len(nums)-k+1; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			tuple = append(tuple, nums[i])
			res = kSum(k-1, i+1, target-int64(nums[i]), nums, res, tuple)
			tuple = tuple[:len(tuple)-1]
		}
	} else {
		l, r := start, len(nums)-1
		for l < r {
			var sum int64 = int64(nums[l]) + int64(nums[r])
			if sum < target {
				l += 1
			} else if sum > target {
				r -= 1
			} else {
				res = append(res, []int{tuple[0], tuple[1], nums[l], nums[r]})
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}
	return res
}
