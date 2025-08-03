package logs

// Application identifies the application emitting the given log.
func Application(log string) string {

	myRuen := []rune(log)

	var output string

	for i := 0; i < len(log); i++ {
		if myRuen[i] == 10071 {
			output = "recommendation"
			return output
		} else if myRuen[i] == rune('🔍') {
			output = "search"
			return output
		} else if myRuen[i] == rune('☀') {
			output = "weather"
			return output
		} else {
			output = "default"
			//fmt.Println(log)
			//fmt.Println(log[0])
		}

	}

	return output
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	myRuen := []rune(log)

	var output string

	for i := range myRuen {
		if myRuen[i] == oldRune {
			myRuen[i] = newRune
		}
	}

	output = string(myRuen)

	return output
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	myRuen := []rune(log)

	return len(myRuen) <= limit
}
