package main

import (
	p "fmt"
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
