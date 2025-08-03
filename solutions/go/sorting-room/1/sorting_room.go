package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	num := nb.Number()
	num2 := float64(num)

	return fmt.Sprintf("This is a box containing the number %.1f", num2)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancy, ok := fnb.(FancyNumber)

	if ok {
		fancyNum, err := strconv.Atoi(fancy.Value())
		if err == nil {
			return fancyNum
		}
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)

	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	numInt, okInt := i.(int)
	if okInt {
		return DescribeNumber(float64(numInt))
	}
	numFloat64, okFloat64 := i.(float64)
	if okFloat64 {
		return DescribeNumber(numFloat64)
	}
	numNumberBox, okNumberBox := i.(NumberBox)
	if okNumberBox {
		return DescribeNumberBox(numNumberBox)
	}
	numFancyNumberBox, okFancyNumberBox := i.(FancyNumberBox)
	if okFancyNumberBox {
		return DescribeFancyNumberBox(numFancyNumberBox)
	}

	return "Return to sender"
}
