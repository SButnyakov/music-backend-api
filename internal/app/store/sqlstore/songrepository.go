package sqlstore

import (
	"database/sql"

	"github.com/SButnyakov/music-backend-api/internal/app/model"
	"github.com/SButnyakov/music-backend-api/internal/app/store"
)

var ()

// Репозиторий песен
type SongRepository struct {
	store *Store
}

// Вставить инфу о новой песне в бд
func (r *SongRepository) Create(s *model.Song) error {
	if err := s.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO songs (artist, name, album) OUTPUT Inserted.id VALUES (@p1, @p2, @p3)",
		s.Artist,
		s.Name,
		s.Album,
	).Scan(&s.ID)

}

// Найти песню по названию
func (r *SongRepository) FindByName(name string) (*model.Song, error) {
	s := &model.Song{}
	if err := r.store.db.QueryRow(
		"SELECT id, artist, name, album FROM songs WHERE name = @p1",
		name,
	).Scan(
		&s.ID,
		&s.Artist,
		&s.Name,
		&s.Album,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return s, nil
}

// Удалить песню
func (r *SongRepository) Delete(artist, name string) (int, error) {
	println(artist)
	println(name)
	id := 0
	if err := r.store.db.QueryRow(
		"DELETE FROM songs OUTPUT Deleted.id WHERE artist = @p1 AND name = @p2",
		artist,
		name,
	).Scan(
		&id,
	); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}

		return 0, err
	}

	return id, nil
}
