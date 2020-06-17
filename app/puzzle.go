package app

import (
    "math/rand"
    "regexp"
    "strconv"
    "strings"
    "time"

    "github.com/LukasJoswiak/wordsearch/models"
)

const (
    min = 1000000
    max = 9999999
)

// Regular expression for transforming a puzzle from what the user input into
// a form ready for database insertion. Replaces end of line with a comma.
var re = regexp.MustCompile(`([ ]*\r?\n)|([ ]*$)`)

var xDir = [...]int{0, 1, 1, 1, 0, -1, -1, -1}
var yDir = [...]int{-1, -1, 0, 1, 1, 1, 0, -1}

func (app *App) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return nil, err
    } else if puzzle == nil {
        return nil, nil
    }
    puzzle.URL = url

    return puzzle, nil
}

func (app *App) GetPuzzleByViewUrl(url string) (*models.Puzzle, error) {
    puzzle, err := app.Database.GetPuzzleByViewUrl(url)
    if err != nil {
        return nil, err
    } else if puzzle == nil {
        return nil, nil
    }

    return puzzle, nil
}

func (app *App) GetFormattedPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.GetPuzzle(url)
    if err != nil {
        return nil, err
    }
    puzzle.Data = strings.Replace(puzzle.Data, ",", "\n", -1)

    return puzzle, nil
}

// Given a puzzle as a string, sanitizes it and returns a copy ready for
// insertion into the database.
func sanitizeBody(body string) string {
    body = re.ReplaceAllString(body, ",")
    body = strings.ToLower(body)
    // Remove comma at end of body.
    if body[len(body) - 1] == ',' {
        body = body[:len(body) - 1]
    }
    return body
}

func (app *App) CreatePuzzle(body string, puzzleType int) (string, error) {
    rand.Seed(time.Now().UnixNano())
    url := strconv.Itoa(rand.Intn(max - min) + min)
    viewUrl := strconv.Itoa(rand.Intn(max - min) + min)
    body = sanitizeBody(body)

    puzzle := &models.Puzzle{
        URL: url,
        ViewURL: viewUrl,
        Data: body,
        Type: puzzleType,
    }
    err := app.Database.CreatePuzzle(puzzle)
    if err != nil {
        return "", err
    }

    return url, nil
}

func (app *App) UpdatePuzzle(url string, body string) error {
    body = sanitizeBody(body)
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return err
    }
    puzzle.Data = body

    err = app.Database.UpdatePuzzle(puzzle)
    if err != nil {
        return err
    }

    return nil
}

func (app *App) ClonePuzzle(url string) (string, error) {
    puzzle, err := app.GetPuzzle(url)
    if err != nil {
        return "", err
    }

    cloneUrl, err := app.CreatePuzzle(puzzle.Data, 2)
    if err != nil {
        return "", err
    }

    return cloneUrl, nil
}

func (app *App) SolvePuzzle(puzzle *models.Puzzle, words *models.Words) *models.SolvedPuzzle {
    puzzleArray := puzzle.ToArray()

    solvedPuzzle := &models.SolvedPuzzle{}
    solvedPuzzle.URL = puzzle.URL

    // Populate each coordinate with the puzzle letter.
    for i := range puzzleArray {
        solvedPuzzle.Locations = append(solvedPuzzle.Locations, []models.Location{})

        for j := range puzzleArray[i] {
            solvedPuzzle.Locations[i] = append(solvedPuzzle.Locations[i], models.Location{})
            solvedPuzzle.Locations[i][j] = models.Location{
                Char: string(puzzleArray[i][j]),
                Coordinate: models.Coordinate{
                    X: j,
                    Y: i,
                },
                Words: []models.Word{},
            }
        }
    }

    // Build up a map of each character to the locations it appears in the puzzle.
    letterMap := letterMap(puzzleArray)

    for index, word := range words.Words {
        startChar := rune(word.Word[0])
        positions := letterMap[startChar]
        found := false

        // Start search from each location first character in word shows up.
        for _, coordinate := range positions {
            xOrig := coordinate.X
            yOrig := coordinate.Y

            // Search in each direction around the starting character.
            for i := 0; i < len(xDir); i++ {
                x := xOrig + xDir[i]
                y := yOrig + yDir[i]

                var j int

                // Search in the selected direction for the length of the word.
                for j = 1; j < len(word.Word); j++ {
                    if x < 0 || y < 0 || y >= len(puzzleArray) || x >= len(puzzleArray[y]) {
                        break
                    }

                    char := puzzleArray[y][x]
                    if char != word.Word[j] {
                        break
                    }

                    x = x + xDir[i]
                    y = y + yDir[i]
                }

                if j == len(word.Word) {
                    // Found word. Add word to each coordinate it appears at in
                    // solved puzzle.
                    found = true
                    x = xOrig
                    y = yOrig
                    for j = 0; j < len(word.Word); j++ {
                        solvedPuzzle.Locations[y][x].Words = append(solvedPuzzle.Locations[y][x].Words, word)

                        x = x + xDir[i]
                        y = y + yDir[i]
                    }
                }
            }
        }

        if found {
            words.Words[index].Exists = true
        }
    }

    return solvedPuzzle
}

// Given an array representation of a puzzle, creates and returns a mapping
// of each character (rune) in the puzzle to an array of Coordinates it appears
// at.
func letterMap(puzzle []string) map[rune][]models.Coordinate {
    m := make(map[rune][]models.Coordinate)

    for y, row := range puzzle {
        for x, char := range row {
            if m[char] == nil {
                m[char] = []models.Coordinate{}
            }
            coordinate := models.Coordinate{
                X: x,
                Y: y,
            }
            m[char] = append(m[char], coordinate)
        }
    }

    return m
}
