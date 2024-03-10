package services

type lifecycles int

const (
	Transient lifecycles = iota
	Singleton
	Scoped
)
