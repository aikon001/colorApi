package color

import (
	"fmt"
	"net/http"
)

type Color struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Hexadecimal string `json:"hexadecimal"`
	R           int    `json:"R"`
	G           int    `json:"G"`
	B           int    `json:"B"`
}

type ColorList struct {
	Color []Color `json:"colors"`
}

func (c *Color) Bind(r *http.Request) error {
	if c.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*ColorList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Color) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
