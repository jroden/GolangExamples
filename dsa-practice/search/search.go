package search

import "errors"

var errFooNFound error = errors.New("value not found in slice")

func linearsearch(arr []int, target int) (int, error) {
	for i, val := range arr {
		if val == target {
			return i, nil
		}
	}
	return 0, errFooNFound
}

func binarysearch(arr []int, target int, low int, high int) (int, error) {
	if high >= low {
		mid := low + ((high - low) / 2)
		if arr[mid] == target {
			return mid, nil
		} else if arr[mid] > target {
			return binarysearch(arr, target, low, mid-1)
		}
		return binarysearch(arr, target, mid+1, high)
	}
	return 0, errFooNFound
}
