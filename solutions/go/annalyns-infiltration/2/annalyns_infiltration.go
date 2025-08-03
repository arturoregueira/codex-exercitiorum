package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var attackSuc bool
	if knightIsAwake {
		attackSuc = false
	} else {
		attackSuc = true
	}
	return attackSuc
	//panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var spySuc bool
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		spySuc = true
	} else {
		spySuc = false
	}
	return spySuc
	//panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	signalSuc := false
	if !archerIsAwake && prisonerIsAwake {
		signalSuc = true
		return signalSuc
	}
	return signalSuc
	//panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	freePrisonerSuc := false
	if !archerIsAwake && petDogIsPresent {
		freePrisonerSuc = true
		return freePrisonerSuc
	}

	if !knightIsAwake && !archerIsAwake && prisonerIsAwake {
		freePrisonerSuc = true
		return freePrisonerSuc
	}

	return freePrisonerSuc

	//panic("Please implement the CanFreePrisoner() function")
}
