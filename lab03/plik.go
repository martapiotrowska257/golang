package main

import (
	"time"
)

type Plik struct {
	name string
	path string
	data[] byte
	createdAt time.Time
	modifiedAt time.Time
}

// --- Metody z FileSystemItem ---

func (p Plik) Name() string {
	return p.name
}

func (p Plik) Path() string {
	return p.path
}

func (p Plik) Size() int64 {
	return int64(len(p.data))
}

func (p Plik) CreatedAt() time.Time {
	return p.createdAt
}

func (p Plik) ModifiedAt() time.Time {
	return p.modifiedAt
}

// --- Metody z Readable ---

func (p Plik) Read(b []byte) (n int, err error) {
	if len(p.data) == 0 {
		return 0, ErrItemNotFound
	}
	n = copy(b, p.data)
	return n, nil
}

// --- Metody z Writable ---

func (p *Plik) Write(b []byte) (n int, err error) {
	if p == nil {
		return 0, ErrNilReference
	}
	p.data = make([]byte, len(b))
	n = copy(p.data, b)
	p.modifiedAt = time.Now()
	return n, nil
}
