package main

import (
	"time"
)

type VirtualFileSystem struct {
	root *Katalog
}

func NewVirtualFileSystem() *VirtualFileSystem {
    root := &Katalog{
        name:       "",
        path:       "/",
        items:      []FileSystemItem{},
        createdAt:  time.Now(),
        modifiedAt: time.Now(),
    }
    return &VirtualFileSystem{root: root}
}

func (vfs *VirtualFileSystem) CreateFile(name string, path string, data []byte) error {
	file := &Plik{
		name:       name,
		path:       path + name,
		data:       data,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(file)
}

func (vfs *VirtualFileSystem) CreateReadFile(name string, path string, data []byte) error {
	file := &PlikDoOdczytu{
		name:       name,
		path:       path + name,
		data:       data,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(file)
}

func (vfs *VirtualFileSystem) CreateDirectory(name string, path string) error {
	directory := &Katalog{
		name:       name,
		path:		path + name,
		items:      []FileSystemItem{},
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(directory)
}

func (vfs *VirtualFileSystem) CreateSymLink(name string, path string, target FileSystemItem) error {
	symlink := &SymLink{
		name:       name,
		path:		path + name,
		target:     target,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(symlink)
}
func (vfs *VirtualFileSystem) FindItem(name string) (FileSystemItem, error) {
	for _, item := range vfs.root.Items() {
		if item.Name() == name {
			return item, nil
		}
	}
	return nil, ErrItemNotFound
}

func (vfs *VirtualFileSystem) ReadFile(name string) ([]byte, error) {
	item, err := vfs.FindItem(name)
	if err != nil {
		return nil, err
	}

	readable, ok := item.(Readable)
	if !ok {
		return nil, ErrPermissionDenied
	}

	buffer := make([]byte, item.Size())
	_, err = readable.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}

func (vfs *VirtualFileSystem) WriteFile(name string, data []byte) error {
	item, err := vfs.FindItem(name)
	if err != nil {
		return err
	}

	writable, ok := item.(Writable)
	if !ok {
		return ErrPermissionDenied
	}

	_, err = writable.Write(data)
	return err
}

func (vfs *VirtualFileSystem) DeleteItem(name string) error {
	return vfs.root.RemoveItem(name)
}