package abstract_factory

type IFile interface {
	Read()
}

type IFolder interface {
	CreateFile() IFile
}
