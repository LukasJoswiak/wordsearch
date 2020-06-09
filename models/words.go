package models

type Words struct {
    PuzzleID int
    Words    []Word
}

type Word struct {
    ID      int
    Word    string
}
