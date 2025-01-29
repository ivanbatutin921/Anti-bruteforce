package interfaceadapters

type IServer interface {
	Start() error
	Stop() error
}
