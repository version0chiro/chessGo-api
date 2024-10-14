package game

import (
	"fmt"
)

const (
	emptySquare = ""
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsValidMove checks if a move from (start_row, start_col) to (end_row, end_col) is valid
func IsValidMove(startRow, startCol, endRow, endCol int, board [][]string) bool {

	// Check if the starting position is valid
	if startRow < 0 || startRow >= 8 || startCol < 0 || startCol >= 8 {
		fmt.Println("Invalid start position")
		return false
	}
	if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
		fmt.Println("Invalid end position")
		return false
	}
	if board[startRow][startCol] == emptySquare {
		fmt.Println("Start position is empty")
		return false
	}

	// Here, you would need to implement the logic to validate the specific piece's move
	// This is a simplified example; you'd need to check the type of piece and its move rules
	piece := board[startRow][startCol]
	fmt.Println("Piece: ", piece)
	fmt.Println("End position: ", endRow, endCol)
	fmt.Println("Start position: ", startRow, startCol)

	// Add piece-specific movement logic here
	// Add piece-specific movement logic here
	switch piece {
	case "P": // White Pawn
		// Single move forward
		if startCol == endCol && endRow == startRow-1 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
		// Double move forward on first move
		if startCol == endCol && endRow == startRow-2 && startRow == 6 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
		// Capture diagonally
		if endRow == startRow-1 && (endCol == startCol-1 || endCol == startCol+1) && board[endRow][endCol] != emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
	case "p": // Black Pawn
		// Single move forward
		if startCol == endCol && endRow == startRow+1 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
		// Double move forward on first move
		if startCol == endCol && endRow == startRow+2 && startRow == 1 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
		// Capture diagonally
		if endRow == startRow+1 && (endCol == startCol-1 || endCol == startCol+1) && board[endRow][endCol] != emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}

	case "R", "r": // Rook
		// Vertical or horizontal move
		if startCol == endCol || startRow == endRow {
			if isPathClear(startRow, startCol, endRow, endCol, board) {
				fmt.Println("Valid rook move")
				return true
			}
		}

	case "B", "b": // Bishop
		// Diagonal move
		if abs(startRow-endRow) == abs(startCol-endCol) {
			if isPathClear(startRow, startCol, endRow, endCol, board) {
				fmt.Println("Valid bishop move")
				return true
			}
		}

	case "Q", "q": // Queen
		// Queen can move like both a rook and a bishop
		if (startCol == endCol || startRow == endRow) || (abs(startRow-endRow) == abs(startCol-endCol)) {
			if isPathClear(startRow, startCol, endRow, endCol, board) {
				fmt.Println("Valid queen move")
				return true
			}
		}

	case "K", "k": // King
		// One square in any direction
		if abs(startRow-endRow) <= 1 && abs(startCol-endCol) <= 1 {
			fmt.Println("Valid king move")
			return true
		}

	case "N", "n": // Knight
		// L-shaped move: 2+1 or 1+2 (no path checking for knights since they jump)
		if (abs(startRow-endRow) == 2 && abs(startCol-endCol) == 1) || (abs(startRow-endRow) == 1 && abs(startCol-endCol) == 2) {
			fmt.Println("Valid knight move")
			return true
		}
	}

	fmt.Println("Invalid move")
	return false
}

// isPathClear checks if there are any pieces in the way between the start and end positions
func isPathClear(startRow, startCol, endRow, endCol int, board [][]string) bool {
	rowStep := 0
	colStep := 0

	// Determine the step direction
	if startRow < endRow {
		rowStep = 1
	} else if startRow > endRow {
		rowStep = -1
	}
	if startCol < endCol {
		colStep = 1
	} else if startCol > endCol {
		colStep = -1
	}

	// Move through the path and check for obstacles
	currentRow := startRow + rowStep
	currentCol := startCol + colStep

	for currentRow != endRow || currentCol != endCol {
		if board[currentRow][currentCol] != emptySquare {
			fmt.Println("Path is blocked at", currentRow, currentCol)
			return false
		}
		currentRow += rowStep
		currentCol += colStep
	}

	return true
}
