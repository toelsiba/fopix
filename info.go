package fopix

import (
	"encoding/json"
	"errors"
	"image"
	"unicode/utf8"
)

type FontInfo struct {
	Name        string
	Author      string
	Description string
	Size        FontSize
	AnchorPos   image.Point
	TargetChar  Character
	CharSet     []RuneInfo
}

type FontSize struct {
	Dx, Dy int
}

type RuneInfo struct {
	Character Character
	Bitmap    []string
}

type Character rune

func (c *Character) MarshalJSON() ([]byte, error) {
	r := rune(*c)
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
