package store

import "github.com/SButnyakov/music-backend-api/internal/app/model"

type SongRepository interface {
	Create(*model.Song) error
	FindByName(string) (*model.Song, error)
	Delete(string, string) (int, error)
}
