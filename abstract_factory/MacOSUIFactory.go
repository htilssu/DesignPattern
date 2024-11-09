package abstract_factory

type MacOSUIFactory struct {
}

type MacOSButton struct {
}

func (m MacOSButton) Click() {
	println("MacOS Button is clicked")
}

func (m MacOSUIFactory) CreateButton() IButton {
	return &MacOSButton{}
}

type MacOSFolder struct {
}

type MacOSFile struct {
}

func (m MacOSFile) Read() {
	println("MacOS File is read")
}

func (m MacOSFolder) CreateFile() IFile {
	return &MacOSFile{}
}

func (m MacOSUIFactory) CreateFolder() IFolder {
	return &MacOSFolder{}
}
