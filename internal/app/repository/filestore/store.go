package filestore

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
)

type Store struct {
	memStore *memstore.Store
	file     *os.File
	rwBuf    *bufio.ReadWriter
}

func NewStore(fileName string) (*Store, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	ms := memstore.NewStore()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)

	store := Store{
		memStore: ms,
		file:     file,
		rwBuf:    bufio.NewReadWriter(reader, writer),
	}

	_ = store.Restore()

	return &store, nil
}

func (s *Store) Restore() error {
	for {
		line, err := s.rwBuf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		shortURLModel := model.ShortURL{}
		err = json.Unmarshal(line, &shortURLModel)
		if err != nil {
			continue
		}

		s.memStore.ShortURLStorage[shortURLModel.Code] = &shortURLModel
	}

	return nil
}

func (s *Store) Commit(overwrite bool) error {
	if overwrite {
		err := s.resetStoreFile()
		if err != nil {
			return err
		}

		err = s.exportMemory2Buffer()
		if err != nil {
			return err
		}
	}
	return s.rwBuf.Flush()
}

func (s *Store) Close() error {
	return s.file.Close()
}

func (s *Store) Write2Buffer(shortURLModel *model.ShortURL) error {
	bShortURL, err := json.Marshal(shortURLModel)
	if err != nil {
		return err
	}

	_, err = s.rwBuf.Write(bShortURL)
	if err != nil {
		return err
	}

	err = s.rwBuf.WriteByte('\n')
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) exportMemory2Buffer() error {
	for _, shortURL := range s.memStore.ShortURLStorage {
		err := s.Write2Buffer(shortURL)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (s *Store) resetStoreFile() error {
	_, err := s.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	return s.file.Truncate(0)
}

// URL return url repository with file store.
func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
