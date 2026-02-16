package app

type Application struct {
	components []Component
	closers    []func() error
}

func New() *Application {
	return &Application{}
}

func (a *Application) Add(c Component) {
	a.components = append(a.components, c)
}

func (a *Application) AddCloser(fn func() error) {
	a.closers = append(a.closers, fn)
}
