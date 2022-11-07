package main

import (
	"strings"

	"github.com/oneitfarm/uppercase"
	"github.com/wasmcloud/actor-tinygo"
)

func main() {
	me := ActorServerGo{}
	actor.RegisterHandlers(uppercase.UppercaseHandler(&me))
}

type ActorServerGo struct{}

func (e *ActorServerGo) Upper(ctx *actor.Context, req uppercase.UppercaseRequest) (*uppercase.UppercaseResponse, error) {
	r := uppercase.UppercaseResponse{
		Data: strings.ToUpper(req.Data),
	}
	return &r, nil
}
