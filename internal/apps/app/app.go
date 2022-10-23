package app

type App interface {
	Name() string
	Entrypoint() string
	Run() error
}
