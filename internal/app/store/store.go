package store

type Store interface {
	Song() SongRepository
}
