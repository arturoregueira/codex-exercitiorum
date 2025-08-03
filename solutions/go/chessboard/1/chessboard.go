package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	booleOfile, WeOk := cb[file]
	i := 0
	if WeOk {
		for _, b := range booleOfile {

			if b {
				i++
			}
		}
	}

	return i

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	mineCounter := 0
	if rank > 0 && rank < 9 {
		for _, booleOfile := range cb {

			daBool := booleOfile[rank-1]

			if daBool {
				mineCounter++
			}

		}
	}

	return mineCounter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	mineOutput := 0
	for _, sliceOBoole := range cb {
		mineOutput += len(sliceOBoole)
	}

	return mineOutput
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	mineOutput := 0
	for file, _ := range cb {
		mineOutput += CountInFile(cb, file)
	}

	return mineOutput
}
