package main

import (
	"github.com/oneitfarm/uppercase"
	"github.com/wasmcloud/actor-tinygo"
	"github.com/wasmcloud/interfaces/httpserver/tinygo"
)

func main() {
	me := ActorClientGo{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me))
}

type ActorClientGo struct{}

func (e *ActorClientGo) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	provider := uppercase.NewActorUppercaseSender("oneitfarm/actor_server")
	res, err := provider.Upper(ctx, uppercase.UppercaseRequest{
		Data: string(req.Body),
	})
	if err != nil {
		return &httpserver.HttpResponse{
			StatusCode: 500,
			Header:     make(httpserver.HeaderMap, 0),
			Body:       []byte(err.Error()),
		}, nil
	}
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte(res.Data),
	}
	return &r, nil
}
