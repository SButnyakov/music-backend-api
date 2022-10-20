package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Структура песни
type Song struct {
	ID     int    `json:"id"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
	Album  string `json:"album"`
}

func (s *Song) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Artist, validation.Required),
		validation.Field(&s.Name, validation.Required))
}
