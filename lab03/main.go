package main

import (
	"fmt"
)

func main() {
	vfs := NewVirtualFileSystem()                                              // tworzenie vfs
	vfs.CreateFile("vfsPlik.txt", "/", []byte("Hello World!"))                 // tworzenie pliku
	vfs.CreateDirectory("vfsKatalog", "/")                                     // tworzenie katalogu
	vfs.CreateReadFile("vfsReadOnlyPlik.txt", "/", []byte("Plik tylko do odczytu")) // tworzenie pliku tylko do odczytu

	// find
	fmt.Println("")
	fmt.Println("FIND")
	vfsFind, err := vfs.FindItem("vfsPlik.txt")
	if err == nil {
		fmt.Println("Znaleziono plik:", vfsFind.Name())
	} else {
		fmt.Println("Error:", err)
	}

	// tworzenie symlink
	vfs.CreateSymLink("vfsSymlink", "/", vfsFind)

	// find err
	fmt.Println("")
	fmt.Println("FIND ERR")
	vfsFindErr, err := vfs.FindItem("vfsNonexisting.txt")
	if err == nil {
		fmt.Println("Znaleziono plik:", vfsFindErr.Name())
	} else {
		fmt.Println("Error:", err)
	}

	// read plik
	fmt.Println("")
	fmt.Println("READ PLIK")
	vfsRead1, err := vfs.ReadFile("vfsPlik.txt")
	if err == nil {
		fmt.Println(string(vfsRead1))
	} else {
		fmt.Println("Error:", err)
	}

	// write plik
	fmt.Println("")
	fmt.Println("WRITE PLIK")
	vfs.WriteFile("vfsPlik.txt", []byte("No, hello universe!"))
	vfsRead2, err := vfs.ReadFile("vfsPlik.txt")
	if err == nil {
		fmt.Println(string(vfsRead2))
	} else {
		fmt.Println("Error:", err)
	}

	// read readonly
	fmt.Println("")
	fmt.Println("READ READONLY")
	vfsReadOnly1, err := vfs.ReadFile("vfsReadOnlyPlik.txt")
	if err == nil {
		fmt.Println(string(vfsReadOnly1))
	} else {
		fmt.Println("Error:", err)
	}

	// write readonly
	fmt.Println("")
	fmt.Println("WRITE READONLY")
	vfs.WriteFile("vfsReadOnlyPlik.txt", []byte("ReadOnly")) // nie nadpisze się, bo jest readonly
	vfsReadOnly2, err := vfs.ReadFile("vfsReadOnlyPlik.txt")
	if err == nil {
		fmt.Println(string(vfsReadOnly2))
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("")
	fmt.Println("Zawartość katalogu root:") // wyświetlanie plików
	for _, item := range vfs.root.Items() {
		fmt.Println(item.Path())
	}

	fmt.Println("")
	fmt.Println("PLIK:")
	plikName := vfs.root.items[0].Name()
	plikPath := vfs.root.items[0].Path()
	plikSize := vfs.root.items[0].Size()
	plikCreatedAt := vfs.root.items[0].CreatedAt()
	plikModifiedAt := vfs.root.items[0].ModifiedAt()
	fmt.Println(plikName)
	fmt.Println(plikPath)
	fmt.Printf("%d bajtów\n", plikSize)
	fmt.Println(plikCreatedAt)
	fmt.Println(plikModifiedAt)

	// usuwanie plik
	fmt.Println("")
	fmt.Println("USUWANIE PLIK")
	vfs.root.RemoveItem("vfsPlik.txt")

	fmt.Println("")
	fmt.Println("Zawartość katalogu root:")
	for _, item := range vfs.root.Items() {
		fmt.Println(item.Path())
	}

	fmt.Println("")
	fmt.Println("DOWIĄZANIE SYMBOLICZNE:")
	symlinkName := vfs.root.items[2].Name()
	symlinkPath := vfs.root.items[2].Path()
	symlinkSize := vfs.root.items[2].Size()
	symlinkCreatedAt := vfs.root.items[2].CreatedAt()
	symlinkModifiedAt := vfs.root.items[2].ModifiedAt()
	fmt.Println(symlinkName)
	fmt.Println(symlinkPath)
	fmt.Printf("%d bajtów\n", symlinkSize)
	fmt.Println(symlinkCreatedAt)
	fmt.Println(symlinkModifiedAt)
}