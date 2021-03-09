package models

import (
	"fmt"
	"net/http"
)

type Color struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Hexadecimal string `json:"hexadecimal"`
	R           int    `json:"r"`
	G           int    `json:"g"`
	B           int    `json:"b"`
}

type ColorList struct {
	Colors []Color `json:"colors"`
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
