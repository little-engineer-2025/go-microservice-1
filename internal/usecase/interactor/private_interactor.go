package interactor

import "github.com/avisiedo/go-microservice-1/internal/interface/interactor"

type private struct{}

func NewPrivate() interactor.Private {
	return newPrivate()
}

func newPrivate() *private {
	return &private{}
}

// TODO Add the implementation below
