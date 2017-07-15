package binarysearch

import "fmt"

const testVersion = 1

func SearchInts(ints []int, find int) int {
	first, lastIdx := 0, len(ints)-1
	for first <= lastIdx {
		middle := (first + lastIdx) / 2
		if find <= ints[middle] {
			lastIdx = middle - 1
		} else {
			first = middle + 1
		}
	}
	return first
}

func Message(slice []int, find int) string {
	idx := SearchInts(slice, find)

	atStart, atEnd := idx == 0, idx == len(slice)-1
	beyondEnd := idx == len(slice)
	found := !beyondEnd && slice[idx] == find

	switch {
	case 0 == len(slice):
		return "slice has no values"
	case found && atStart:
		return fmt.Sprintf("%d found at beginning of slice", find)
	case found && atEnd:
		return fmt.Sprintf("%d found at end of slice", find)
	case found:
		return fmt.Sprintf("%d found at index %d", find, idx)
	case !found && atStart:
		return fmt.Sprintf("%d < all values", find)
	case !found && beyondEnd:
		return fmt.Sprintf("%d > all %d values", find, len(slice))
	case !found:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
			find, slice[idx-1], idx-1, slice[idx], idx)
	default:
		return ""
	}
}
