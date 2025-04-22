package main

import (
	"time"
)

type Katalog struct {
	name string
	path string
	items []FileSystemItem
	createdAt time.Time
	modifiedAt time.Time
}

// --- Metody z Directory ---

// --- Metody z FileSystemItem ---
func (k Katalog) Name() string {
	return k.name
}

func (k Katalog) Path() string {
	return k.path
}

func (k Katalog) Size() int64 {
	return int64(len(k.items))
}

func (k Katalog) CreatedAt() time.Time {
	return k.createdAt
}

func (k Katalog) ModifiedAt() time.Time {
	return k.modifiedAt
}

func (k *Katalog) AddItem(item FileSystemItem) error {
    // check for duplicates
    for _, existingItem := range k.items {
        if existingItem.Name() == item.Name() {
            return ErrItemExists
        }
    }
    
    k.items = append(k.items, item)
    k.modifiedAt = time.Now()
    return nil
}

func (k *Katalog) RemoveItem(name string) error {
	for i, item := range k.items {
		if item.Name() == name {
			k.items = append(k.items[:i], k.items[i+1:]...)
			k.modifiedAt = time.Now()
			return nil
		}
	}
	return ErrItemNotFound
}

func (k *Katalog) Items() []FileSystemItem {
	return k.items
}