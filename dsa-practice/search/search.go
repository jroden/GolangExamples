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

func jumpsearch(arr []int, target int, jump int) (int, error) {
	// check error conditions
	if arr[len(arr)-1] < target {
		return 0, errFooNFound
	} else if jump <= 0 {
		return 0, errors.New("Please specify a valid jump value")
	}
	// jump through array, checking if value is less than target
	for index, previndex := 0, 0; index < len(arr); index = index + jump {
		if arr[index] == target {
			return index, nil
		} else if arr[index] > target {
			// if value is less than target, iterate starting from previndex
			for subindex := previndex; subindex < index; subindex++ {
				if arr[subindex] == target {
					return subindex, nil
				}
			}
			return 0, errFooNFound
		}
		previndex = index
	}
	// iterate through remaining values if next jump is past limit
	remainder := len(arr) % jump
	for index := len(arr) - remainder; index < len(arr); index++ {
		if arr[index] == target {
			return index, nil
		}
	}
	return 0, errFooNFound
}
