package file

import (
	"bytes"
	"time"
)

type File struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      bytes.Buffer
}

func New(created time.Time, data bytes.Buffer) *File {
	return &File{
		Name:      "",
		CreatedAt: created,
		UpdatedAt: created,
		Data:      data,
	}
}

type Store interface {
	Save(file *File) (string, error)
	GetByName(name string) (*File, error)
	ListAll() []*File
}

type InMemoryStore struct {
}

func (i InMemoryStore) Save(file *File) (string, error) {
	panic("implement me")
}

func (i InMemoryStore) GetByName(name string) (*File, error) {
	panic("implement me")
}

func (i InMemoryStore) ListAll() []*File {
	panic("implement me")
}
