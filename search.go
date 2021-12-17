package binary_search

import "errors"

func BinarySearch(orderedArray []int, target int) (int, error) {
	leftPosition := 0
	rightPosition := len(orderedArray) - 1

	for leftPosition <= rightPosition {
		if orderedArray[leftPosition] == target {
			return leftPosition, nil
		}

		if orderedArray[rightPosition] == target {
			return rightPosition, nil
		}

		middlePosition := (leftPosition + rightPosition) / 2
		if orderedArray[middlePosition] == target {
			return middlePosition, nil
		}

		switch {
		case target < orderedArray[middlePosition]:
			rightPosition = middlePosition - 1
		default:
			leftPosition = middlePosition + 1
		}
	}

	return 0, errors.New("not found")
}
