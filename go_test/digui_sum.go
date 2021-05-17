package main

func main() {
	arr := []int{1, 7, 2, 4, 3, 8, 4, 5, 7, 6, 7, 8}
	println(sum(arr))
	println(num(arr))
	println(max(arr))
}

func sum(arr []int) (x int) {
	if len(arr) == 1 {
		return arr[0]
	} else {
		x, arr = arr[0], arr[1:]
		return x + sum(arr)
	}
}

func num(arr []int) (n int) {
	if len(arr) == 1 {
		return 1
	} else {
		arr = arr[1:]
		return 1 + num(arr)
	}
}

func max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		if arr[0] < arr[1] {
			arr = arr[1:]
		} else {
			arr = append(arr[:1], arr[2:]...)
		}
		return max(arr)
	}
}
