package quicksort

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// in order to generate really random numbers
	rand.Seed(time.Now().UnixNano())
}

// newArray creates a array that contains a random set of numbers for test purposes
func newArray(max byte) (numbers array2Sort, err error) {
	var zero byte
	if max <= zero {
		err = errors.New("incorrect max")
		return
	}

	for i := zero; i < 32; i++ {
		numbers[i] = rand.Int31n(int32(max))
	}

	return
}

func describe(stageName string, numbers *array2Sort, wall int, pivot int32, index int) {
	fmt.Println(stageName)
	fmt.Println(*numbers)
	fmt.Println("wall is", wall)
	fmt.Println("pivot is", pivot)
	fmt.Println("index is", index)
	fmt.Println()
}

// sort sorts numbers in the array
func sort(numbers *array2Sort) {
	defer trackTime(time.Now())

	maxIndex := 31
	steps := 0
	var minNumber int32
	var minIndex int
	var limit byte

	fmt.Println(numbers)

	for i := 0; i <= maxIndex; i++ {

		if i == 0 {
			minNumber = numbers[0]
		} else {
			if numbers[i] < minNumber {
				minNumber = numbers[i]
			}
		}
	}

	for {
		pivot := rand.Int31n(int32(maxIndex))
		//pivot := int32(minIndex) + rand.Int31n(int32(maxIndex-minIndex))
		wall := 0

		for i := 0; i <= maxIndex; i++ {
			//describe("=== new iteration:", numbers, wall, pivot, i)

			switch {
			// check if the current number is smaller than the pivot
			case numbers[i] < pivot:

				if wall != i {
					numbers[wall], numbers[i] = numbers[i], numbers[wall]

					//describe("  -- the number is smaller than the pivot", numbers, wall, pivot, i)
				} else {
					//fmt.Println("  -- no need to exchange")
				}

				wall++
				//fmt.Println("  -- move the wall:", wall)

			default:
				//fmt.Println("  -- the number is bigger than the pivot")
				//fmt.Println()
			}
		}

		if numbers[0] == minNumber {
			var begin = 1
			if minIndex != 0 {
				begin = minIndex
			}
			//fmt.Println("begin", begin)
			for i := begin; i <= maxIndex; i++ {
				if numbers[i] < numbers[i-1] {
					minIndex = i-1
					break
				} else {
					minIndex = i
				}
			}
			//fmt.Println(numbers)
			//fmt.Println(minIndex)
		}

		if minIndex == maxIndex {
			fmt.Println("Sorted!, steps:", steps)
			//fmt.Println("Minimum is", minNumber)
			fmt.Println(numbers)
			break
		}

		limit++
		if limit == 200 {
			fmt.Println("too long")
			break
		}
		steps++
	}
}

func trackTime(start time.Time) {
	elapsed := time.Since(start)

	fmt.Println("Time elapsed:", elapsed.Nanoseconds(), "ns")
}
