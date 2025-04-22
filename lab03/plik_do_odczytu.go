package main

import (
	"time"
)

type PlikDoOdczytu struct {
	name string
	path string
	data[] byte
	createdAt time.Time
	modifiedAt time.Time
}

// --- Metody z FileSystemItem ---

func (pd PlikDoOdczytu) Name() string {
	return pd.name
}

func (pd PlikDoOdczytu) Path() string {
	return pd.path
}

func (pd PlikDoOdczytu) Size() int64 {
	return int64(len(pd.data))
}

func (pd PlikDoOdczytu) CreatedAt() time.Time {
	return pd.createdAt
}

func (pd PlikDoOdczytu) ModifiedAt() time.Time {
	return pd.modifiedAt
}

// --- Metody z Readable ---

func (pd PlikDoOdczytu) Read(b []byte) (n int, err error) {
	if len(pd.data) == 0 {
		return 0, ErrItemNotFound
	}
	n = copy(b, pd.data)
	return n, nil
}