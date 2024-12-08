package utils

import "fmt"

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Find[T comparable](elems []T, v T) int {
	for i, s := range elems {
		if v == s {
			return i
		}
	}
	return -1
}

func FindIf[T any](elems []T, f func(elem T) bool) int {
	for i, s := range elems {
		if f(s) {
			return i
		}
	}
	return -1
}

func Filter[T any](elems []T, f func(T) bool) []T {
	var res []T
	for _, s := range elems {
		if f(s) {
			res = append(res, s)
		}
	}
	return res
}

func Map[T1, T2 any](elems []T1, f func(T1) T2) []T2 {
	var res []T2
	for _, s := range elems {
		res = append(res, f(s))
	}
	return res
}

func Reduce[T any, U any](elems []T, f func(U, T) U, init U) U {
	res := init
	for _, s := range elems {
		res = f(res, s)
	}
	return res
}

func Reverse[T any](elems []T) []T {
	res := make([]T, len(elems))
	for i, s := range elems {
		res[len(elems)-1-i] = s
	}
	return res
}

func Sum[T any](elems []T, f func(T) int) int {
	res := 0
	for _, s := range elems {
		res += f(s)
	}
	return res
}

func Max[T any](elems []T, f func(T) int) int {
	res := f(elems[0])
	for _, s := range elems {
		if f(s) > res {
			res = f(s)
		}
	}
	return res

}

func Min[T any](elems []T, f func(T) int) int {
	res := f(elems[0])
	for _, s := range elems {
		if f(s) < res {
			res = f(s)
		}
	}
	return res
}

func Copy[T any](elems []T) []T {
	copy_ := make([]T, len(elems))
	copy(copy_, elems)
	return copy_
}

func Copy2D[T any](elems [][]T) [][]T {
	copy_ := make([][]T, len(elems))
	for i := range elems {
		copy_[i] = make([]T, len(elems[i]))
		copy(copy_[i], elems[i])
	}
	return copy_
}

func StringsToints(elems []string) [][]int {
	var res [][]int
	for _, s := range elems {
		var row []int
		for _, c := range s {
			row = append(row, int(c))
		}
		res = append(res, row)
	}
	return res
}

func StringsToChars(elems []string) [][]rune {
	var res [][]rune
	for _, s := range elems {
		var row []rune
		for _, c := range s {
			row = append(row, c)
		}
		res = append(res, row)
	}
	return res
}

func StringsToStrings(elems []string) [][]string {
	var res [][]string
	for _, s := range elems {
		var row []string
		for _, c := range s {
			row = append(row, string(c))
		}
		res = append(res, row)
	}
	return res

}

func PrintMap[T any](elems [][]T) {
	for _, elem := range elems {
		for _, e := range elem {
			fmt.Print(e)
		}
		fmt.Println()
	}
}

func initialize2D[T any](row int, col int) [][]T {
	board := make([][]T, row)
	for i := range board {
		board[i] = make([]T, col)
	}
	return board
}
