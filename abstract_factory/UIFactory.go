package abstract_factory

type UIFactory interface {
	CreateButton() IButton
	CreateFolder() IFolder
}
