package app

import (
    "math/rand"
    "regexp"
    "strconv"
    "strings"
    "time"
    "unicode"
    "unicode/utf8"

    "github.com/LukasJoswiak/wordsearch/models"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

// Min and max values to use when generating random URLs.
const (
    min = 1000000000
    max = 9999999999
)

// Regular expression for transforming a puzzle from what the user input into
// a form ready for database insertion. Replaces end of line with a comma.
var re = regexp.MustCompile(`([ ]*\r?\n)|([ ]*$)`)

var xDir = [...]int{0, 1, 1, 1, 0, -1, -1, -1}
var yDir = [...]int{-1, -1, 0, 1, 1, 1, 0, -1}

// Used for normalizing Unicode characters.
var t = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

func isMn(r rune) bool {
    return unicode.Is(unicode.Mn, r)
}

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
    puzzle.ViewURL = url

    return puzzle, nil
}

func (app *App) GetFormattedPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.GetPuzzle(url)
    if err != nil {
        return nil, err
    } else if puzzle == nil {
        return nil, nil
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
    // TODO: Make sure URL isn't a duplicate
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

func (app *App) ClonePuzzle(viewUrl string) (string, error) {
    puzzle, err := app.GetPuzzleByViewUrl(viewUrl)
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
    solvedPuzzle.ViewURL = puzzle.ViewURL

    // Populate each coordinate with the puzzle letter.
    for i := range puzzleArray {
        solvedPuzzle.Locations = append(solvedPuzzle.Locations, []models.Location{})

        // Unicode characters can take up several bytes, causing the index in
        // the range statement to "skip". The solution is to keep track of the
        // index separately.
        index := 0
        for _, c := range puzzleArray[i] {
            solvedPuzzle.Locations[i] = append(solvedPuzzle.Locations[i], models.Location{})

            if index < len(solvedPuzzle.Locations[i]) {
                solvedPuzzle.Locations[i][index] = models.Location{
                    Char: string(c),
                    Coordinate: models.Coordinate{
                        X: index,
                        Y: i,
                    },
                    Words: []models.Word{},
                    Class: "",
                }
                index += 1
            }
        }
    }

    // Build up a map of each character to the locations it appears in the puzzle.
    letterMap := letterMap(puzzleArray)

    for index, word := range words.Words {
        // Normalize first character.
        startChar, _, _ := transform.String(t, string([]rune(word.Word)[0]))
        startRune := rune(startChar[0])
        positions := letterMap[startRune]
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
                for j = 1; j < utf8.RuneCountInString(word.Word); j++ {
                    if x < 0 || y < 0 || y >= len(puzzleArray) || x >= utf8.RuneCountInString(puzzleArray[y]) {
                        break
                    }

                    // Normalize Unicode characters.
                    char, _, _ := transform.String(t, string([]rune(puzzleArray[y])[x]))
                    wordChar, _, _ := transform.String(t, string([]rune(word.Word)[j]))
                    if char != wordChar {
                        break
                    }

                    x = x + xDir[i]
                    y = y + yDir[i]
                }

                if j == utf8.RuneCountInString(word.Word) {
                    // Found word. Add word to each coordinate it appears at in
                    // solved puzzle.
                    found = true
                    x = xOrig
                    y = yOrig
                    for j = 0; j < utf8.RuneCountInString(word.Word); j++ {
                        solvedPuzzle.Locations[y][x].Words = append(solvedPuzzle.Locations[y][x].Words, word)
                        solvedPuzzle.Locations[y][x].Class += word.Word + " "

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
        row, _, _ := transform.String(t, row)
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
