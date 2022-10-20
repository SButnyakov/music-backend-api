package model

import "testing"

func TestSong(t *testing.T) *Song {
	return &Song{
		Artist: "artist",
		Name:   "name",
		Album:  "album",
	}
}
