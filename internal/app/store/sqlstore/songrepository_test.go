package sqlstore_test

import (
	"testing"

	"github.com/SButnyakov/music-backend-api/internal/app/model"
	"github.com/SButnyakov/music-backend-api/internal/app/store"
	"github.com/SButnyakov/music-backend-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestSongRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("songs")

	s := sqlstore.New(db)
	song := model.TestSong(t)
	assert.NoError(t, s.Song().Create(song))
	assert.NotNil(t, song)
}

func TestSongRepository_FindByName(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("songs")

	s := sqlstore.New(db)
	name := "name"
	_, err := s.Song().FindByName(name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	song := model.TestSong(t)
	song.Name = name
	s.Song().Create(song)
	song, err = s.Song().FindByName(name)
	assert.NoError(t, err)
	assert.NotNil(t, song)
}

func TestUserRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("songs")

	s := sqlstore.New(db)
	song := model.TestSong(t)
	id, err := s.Song().Delete(song.Artist, song.Name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	assert.EqualValues(t, 0, id)

	s.Song().Create(song)
	id, err = s.Song().Delete(song.Artist, song.Name)
	assert.NoError(t, err)
	assert.NotEqualValues(t, 0, id)
}
