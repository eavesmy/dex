package dex

import "context"

type Aurora struct {
	*Client
}

func (a *Aurora) Init(ctx context.Context) (aurora *Aurora, err error) {

	a.Client = &Client{
		ctx: ctx,
	}

	aurora = a

	return
}
