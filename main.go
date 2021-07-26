package main

func IndexOfMaxAndMaxElement(numbers []int) (int, int) {
	Index := 0
	MaxElement := numbers[0]
	for index, value := range numbers{
		if MaxElement < value{
		MaxElement = value
		Index = index
		}
	}
	return Index, MaxElement
}

func IndexOfMaxElement(numbers []int) int {
	Index := 0
	MaxElement := numbers[0]
	for index, value := range numbers {
		if MaxElement < value {
			MaxElement = value
			Index = index
		}
	}
	return Index
}

func main() {
IndexOfMaxElement([]int {10, -5, 1, -100, 25})
}
