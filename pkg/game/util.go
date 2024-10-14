package game

import (
	"fmt"
)

const (
	emptySquare = ""
)

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
	switch piece {
	case "P": // Pawn
		if startCol == endCol && endRow == startRow-1 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
		// Add more pawn rules (captures, initial two-square move, etc.)
	case "p": // Black pawn
		if startCol == endCol && endRow == startRow+1 && board[endRow][endCol] == emptySquare {
			fmt.Println("Valid pawn move")
			return true
		}
	}

	// Add rules for other pieces (Rooks, Knights, Bishops, Queens, Kings)

	return false
}
