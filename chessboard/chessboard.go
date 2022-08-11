package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	squares := cb[file]
    var occupiedSquares int
    for _, square := range squares {
        if square {
            occupiedSquares++
        }
    }
	return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var occupiedSquares int
    rankZeroIndex := rank - 1
    for _, file := range cb {
        for j, square := range file {
            if j == rankZeroIndex && square {
                occupiedSquares++
            }
        }
    }
	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares int
    for _, file := range cb {
        for range file {
            squares++
        }
    }
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupiedSquares int
    for file := range cb {
        occupiedSquares += CountInFile(cb, file)
    }
	return occupiedSquares
}
