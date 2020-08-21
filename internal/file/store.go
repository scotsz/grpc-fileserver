package file

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type File struct {
	Data bytes.Buffer
	*Info
}
type Info struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func New(name string, created time.Time, data bytes.Buffer) *File {
	return &File{
		Info: &Info{
			CreatedAt: created,
			UpdatedAt: created,
			Name:      name,
		},
		Data: data,
	}
}

type Store interface {
	Save(file *File) error
	ListAll() []*Info
	ReadFile(name string) (*os.File, error)
}

type systemStore struct {
	files map[string]*Info
	mu    *sync.Mutex
	dir   string
}

func NewSystemStore(dir string) *systemStore {
	return &systemStore{
		files: make(map[string]*Info),
		mu:    &sync.Mutex{},
		dir:   dir,
	}
}

func (s systemStore) Save(file *File) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", s.dir, file.Name))
	if err != nil {
		return err
	}
	_, err = file.Data.WriteTo(f)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.files[file.Name] = file.Info
	s.mu.Unlock()
	return nil
}

func (s systemStore) ReadFile(name string) (*os.File, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s", s.dir, name))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s systemStore) ListAll() []*Info {
	files := make([]*Info, len(s.files))
	s.mu.Lock()
	for _, f := range s.files {
		files = append(files, f)
	}
	s.mu.Unlock()
	return files
}
