package fopix

import (
	"encoding/json"
	"errors"
	"unicode/utf8"
)

type FontInfo struct {
	Name        string     `json:"name"`
	Author      string     `json:"author"`
	Description string     `json:"description"`
	Size        Point      `json:"size"`
	AnchorPos   Point      `json:"anchor-pos"`
	TargetChar  Character  `json:"target-char"`
	CharSet     []RuneInfo `json:"char-set"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type RuneInfo struct {
	Character Character `json:"character"`
	Bitmap    []string  `json:"bitmap"`
}

type Character rune

func (c Character) MarshalJSON() ([]byte, error) {
	r := rune(c)
	if !utf8.ValidRune(r) {
		return nil, errors.New("rune is not valid")
	}
	return json.Marshal(string(r))
}

func (c *Character) UnmarshalJSON(data []byte) error {

	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return errors.New("error rune value")
	}
	if size != len(s) {
		return errors.New("error rune size")
	}

	*c = Character(r)

	return nil
}

type ByCharacter []RuneInfo

func (s ByCharacter) Len() int {
	return len(s)
}
func (s ByCharacter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByCharacter) Less(i, j int) bool {
	return s[i].Character < s[j].Character
}
