package pipeline

import (
	"net/http"
)

type ComponentsContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mwc *ComponentsContext) Error(err error) {
	mwc.error = err
}

func (mwc *ComponentsContext) GetError() error {
	return mwc.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentsContext, next func(*ComponentsContext))
}

type ServicesMiddlwareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
