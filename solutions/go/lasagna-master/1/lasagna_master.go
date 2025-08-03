package lasagna //dam

import "math"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		time = 2
	}

	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(stuff []string) (int, float64) {
	noodles, sauce := 0, 0.0

	for _, v := range stuff {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2

		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func roundTo(num float64, place int) float64 {
	factor := math.Pow(10, float64(place))
	return math.Round(num*factor) / factor
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, servings int) []float64 {
	portionsPerRec := 2.0

	multiplyer := float64(servings) / portionsPerRec

	var scaledQuantities []float64
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, roundTo(v*multiplyer, 2))
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
