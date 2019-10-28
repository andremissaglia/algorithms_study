package binarysearch

import (
	"errors"

	"github.com/andremissaglia/algorithms_study/internal/utils"
)

// ErrNotFound means the desired @key is not in the @items
var ErrNotFound error = errors.New("Not Found")

// BinarySearch looks for the element in @items with the matching @key
func BinarySearch(items []utils.Element, key string) (utils.Element, error) {
	start := 0
	end := len(items) - 1
	for end-start >= 0 {
		middle := (start + end) / 2
		elem := items[middle]
		if elem.Key == key {
			return elem, nil
		}

		if elem.Key < key {
			// search on the right side
			start = middle + 1
		} else {
			// search on the left side
			end = middle - 1
		}
	}
	return utils.Element{}, ErrNotFound
}
