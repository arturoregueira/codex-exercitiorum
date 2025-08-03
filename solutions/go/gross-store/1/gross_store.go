package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {

	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	value, ok := units[unit]

	if ok {
		amount, already_there := bill[item]

		if already_there {
			bill[item] = value + amount
		} else {
			bill[item] = value // Did this quick duble check.
		}

	}

	return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	value, ok := units[unit]
	amount, already_there := bill[item]

	if ok && already_there {

		if value > amount {
			ok = false
		} else if value == amount {
			delete(bill, item)
		} else {
			bill[item] = amount - value
		}

	}

	return ok && already_there
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, already_there := bill[item]

	return amount, already_there
}
