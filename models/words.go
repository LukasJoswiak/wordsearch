package models

type Words struct {
    PuzzleID int
    Words    []Word
}

type Word struct {
    ID      int
    Word    string
    Exists  bool
}

// Form input.
type WordsForm struct {
    Words []WordInput
}

type WordInput struct {
    ExistingWord string
    Word         string
}
