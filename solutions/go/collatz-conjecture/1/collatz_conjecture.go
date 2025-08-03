package collatzconjecture // Solved wityh niave solution

import "errors"

func CollatzConjecture(n int) (int, error) {
	amountToSolution := 0
	if n > 0 {
		for i := 0; n != 1; i++ {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = n*3 + 1
			}
			amountToSolution++
		}

		return amountToSolution, nil
	}

	return 0, errors.New("i hope tis wok")

}
