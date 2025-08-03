package raindrops // Finished niave solution

import "fmt"

var dropNums = []int{3, 5, 7}
var dropWords = []string{"Pling", "Plang", "Plong"}

func Convert(number int) string {
	var rainDrop string
	var foundButADrop bool = false

	for i, v := range dropNums {
		//fmt.Println(k, v)
		if number%v == 0 {
			rainDrop = rainDrop + dropWords[i]
			foundButADrop = true
		}
	}

	if !foundButADrop {
		return fmt.Sprint(number)
	}

	return rainDrop
	//return "This is a test"
}
