package abstract_factory

import "fmt"

type WindowUIFactory struct {
}

func (w WindowUIFactory) CreateButton() IButton {
	return &WindowButton{}
}

func (w WindowUIFactory) CreateFolder() IFolder {
	return &WindowFolder{}
}

type WindowButton struct{}

func (w WindowButton) Click() {
	fmt.Printf("Window Button is clicked")
}

type WindowFolder struct {
}

type WindowFile struct {
}

func (w WindowFile) Read() {
	fmt.Printf("Window File is read")
}

func (w *WindowFolder) CreateFile() IFile {
	return &WindowFile{}
}
