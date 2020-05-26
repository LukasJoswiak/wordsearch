package models

type Puzzle struct {
    URL    string `json:"url"`
    Width  int    `json:"width"`
    Height int    `json:"height"`
    Data   string `json:"data"`
}
