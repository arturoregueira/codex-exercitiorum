package lasagna

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {

	remaining_minutes := OvenTime - actualMinutesInOven

	return remaining_minutes
	//panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	const layerLapse = 2
	prep_time := numberOfLayers * layerLapse
	return prep_time
	//panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	prep_time := PreparationTime(numberOfLayers)
	//remaining_minutes := RemainingOvenTime(actualMinutesInOven)
	time := prep_time + actualMinutesInOven
	return time
	//panic("ElapsedTime not implemented")
}
