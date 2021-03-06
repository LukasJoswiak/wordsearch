package models

import "strings"

type Puzzle struct {
    ID         int    `json:"id"`
    URL        string `json:"url"`
    ViewURL    string `json:"viewUrl"`
    Data       string `json:"data"`
    Type       int    `json:"type"`
}

type SolvedPuzzle struct {
    URL           string
    ViewURL       string
    Locations     [][]Location
}

type Location struct {
    Char string
    Coordinate Coordinate
    Words []Word
    Class string
}

type Coordinate struct {
    X int
    Y int
}

// Returns an array representation of the puzzle, where each index represents
// a row of characters.
func (puzzle *Puzzle) ToArray() []string {
    return strings.Split(puzzle.Data, ",")
}
