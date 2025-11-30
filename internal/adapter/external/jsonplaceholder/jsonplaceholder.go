package jsonplaceholder

import "context"

type jsonPlaceHolder struct{}

type JsonPlaceHolder interface {
	FetchExternal(ctx context.Context) error
}

func NewJsonPlaceHolder() *jsonPlaceHolder {
	return &jsonPlaceHolder{}
}
