package model_test

import (
	"testing"

	"github.com/SButnyakov/music-backend-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestSong_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.Song
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Song {
				return model.TestSong(t)
			},
			isValid: true,
		},
		{
			name: "without album",
			s: func() *model.Song {
				s := model.TestSong(t)
				s.Album = ""
				return s
			},
			isValid: true,
		},
		{
			name: "without artist",
			s: func() *model.Song {
				s := model.TestSong(t)
				s.Artist = ""
				return s
			},
			isValid: false,
		},
		{
			name: "without name",
			s: func() *model.Song {
				s := model.TestSong(t)
				s.Name = ""
				return s
			},
			isValid: false,
		},
		{
			name: "without name and artist",
			s: func() *model.Song {
				s := model.TestSong(t)
				s.Artist = ""
				s.Name = ""
				return s
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
