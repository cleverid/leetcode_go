package utils

// pop from queue. First element
// x, a = a[0], a[1:]

// pop from stack. Last element
// x, a = a[len(a)-1], a[:len(a)-1]

// push
// a = append(a, x)

func PopQueue[T any](arr []T) (*T, []T) {
	if len(arr) == 0 {
		return nil, arr
	}
	return &arr[0], arr[1:]
}

func PopStack[T any](arr []T) (*T, []T) {
	if len(arr) == 0 {
		return nil, arr
	}
	return &arr[len(arr)-1], arr[:len(arr)-1]
}
