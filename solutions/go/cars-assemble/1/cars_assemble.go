package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0

	//panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)

	//panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const COST_PER_ONE = 10000
	const COST_PER_TEN = 95000

	ofTen := carsCount / 10
	ofOne := carsCount % 10
	totalCost := uint((ofTen * COST_PER_TEN) + (ofOne * COST_PER_ONE))
	return totalCost
	//panic("CalculateCost not implemented")
}
