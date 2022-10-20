package sqlstore

import (
	"database/sql"

	"github.com/SButnyakov/music-backend-api/internal/app/store"
	_ "github.com/denisenkom/go-mssqldb"
)

// Структура хранилища
type Store struct {
	db             *sql.DB
	songRepository *SongRepository
}

// Конструктор хранилища
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Возвращает инстанс репозитория
// Нужно, чтобы не было дубликатов репыы
func (s *Store) Song() store.SongRepository {
	if s.songRepository != nil {
		return s.songRepository
	}

	s.songRepository = &SongRepository{
		store: s,
	}

	return s.songRepository
}
